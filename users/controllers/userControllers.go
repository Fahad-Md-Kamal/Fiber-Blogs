package controllers

import (
	"github.com/fahad-md-kamal/fiber-blogs/users/dtos"
	"github.com/gofiber/fiber/v2"
)

func AddUserHandler(c *fiber.Ctx) error {

	var userCreateDto dtos.UserCreateDto

	if err := c.BodyParser(&userCreateDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if errors, ok := userCreateDto.ValidateUserCreateDto(); !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors})
	}

	UserToCreate := userCreateDto.ParseFromDto()
	if err, ok := UserToCreate.GeneratePasswordHash(); !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if message, ok := UserToCreate.ValidateUserExists(); ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": message})
	}

	if err := UserToCreate.Save(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	responseDto := new(dtos.UserResponseDto)
	responseDto.ParseToResponseDto(UserToCreate)

	return c.JSON(fiber.Map{
		"data": responseDto,
	})
}

func GetUsersListHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Yeee!!!, Fiber Project has started",
	})
}
