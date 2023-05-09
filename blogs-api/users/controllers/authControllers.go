package controllers

import (
	"log"
	"strings"

	"github.com/fahad-md-kamal/fiber-blogs/users/dtos"
	"github.com/fahad-md-kamal/fiber-blogs/users/helpers"
	"github.com/fahad-md-kamal/fiber-blogs/users/models"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	loginRequestData := dtos.LoginRequestDto{}
	if err := c.BodyParser(&loginRequestData); err != nil {
		log.Printf("Error parsing Login Request: %s | Error: %s", c.Params("id"), err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Login Data"})
	}

	if errors, ok := loginRequestData.ValidateLoginRequestDto(); !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors})
	}

	userCheckParams := models.UserCheckParams{
		Username: loginRequestData.Username,
		Email:    loginRequestData.Username,
	}
	dbUser, _, exists := models.ValidateUserExistsWithEmailOrUsername(userCheckParams)
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	if !dbUser.IsActive {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "User Inactive"})
	}

	if msg, ok := dbUser.ValidatePasswordHash(loginRequestData.Password); !ok {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": msg})
	}

	tokenString, success := helpers.GenerateJwtToken(dbUser)
	if !success {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Faild to create auth token"})
	}
	userResponseDto := dtos.ParseToLoginResponseDto(tokenString, dbUser)
	return c.JSON(&userResponseDto)
}

func LogoutHandler(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"message": "Missing Authorization header",
		})
	}

	authToken := strings.Split(authHeader, " ")[1]

	// Invalidate the token by adding it to the blacklist
	blacklistedToken := models.BlacklistedTokens{Token: authToken}

	if ok := blacklistedToken.IsTokenBlacklisted(); ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token or token expired",
		})
	}

	if ok := blacklistedToken.BlacklistToken(); !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Couldn't blacklist token",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User logged out",
	})
}
