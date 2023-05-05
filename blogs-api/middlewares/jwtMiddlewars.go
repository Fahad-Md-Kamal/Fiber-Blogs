package middlewares

import (
	"github.com/fahad-md-kamal/fiber-blogs/configs"
	userdtos "github.com/fahad-md-kamal/fiber-blogs/users/dtos"
	usermodels "github.com/fahad-md-kamal/fiber-blogs/users/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func JwtMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization header"})
		}
		tokenString := authHeader[len("Bearer "):]

		blacklistedToken := usermodels.BlacklistedTokens{Token: tokenString}
		if blacklistedToken.IsTokenBlacklisted() {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token or token expired",
			})
		}

		token, err := jwt.ParseWithClaims(tokenString, &userdtos.TokenClaimPayload{},
			func(token *jwt.Token) (interface{}, error) {
				// Get the signing key from your authentication server or config file
				signingKey := []byte(configs.ENVs.JwtSecretKey)
				return signingKey, nil
			})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid JWT signature"})
			}
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid JWT token"})
		}

		// Check if token is valid
		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid JWT token"})
		}

		// Extract claims from JWT token
		claims, ok := token.Claims.(*userdtos.TokenClaimPayload)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid JWT token"})
		}

		var user usermodels.Users
		user.ID = claims.ID
		user.Username = claims.Username
		user.IsSuperuser = claims.IsSuperuser
		user.Email = claims.Email
		c.Locals("user", user)
		return c.Next()
	}
}
