package controllers

import (
	"strconv"

	"github.com/fahad-md-kamal/fiber-blogs/users/dtos"
	"github.com/fahad-md-kamal/fiber-blogs/users/models"
	"github.com/fahad-md-kamal/fiber-blogs/utils"
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

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": responseDto,
	})
}

func GetUsersListHandler(c *fiber.Ctx) error {

	// Parse pagination parameters
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset := (page - 1) * limit

	// Get Users List
	users, totalCount, err := models.GetUsersList(limit, offset)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	// Convert User's list into response Dtos
	userDtos := dtos.ParseUsersListToResponseDto(&users)

	// Get Paginated Response
	pagination := utils.Paginate(int(totalCount), limit, page, userDtos)
	return c.JSON(pagination)
}

func GetUserDetailHandler(c *fiber.Ctx) error {
	userId, err := strconv.ParseUint(c.Params("id"), 10, 0)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Invalid User Id",
		})
	}

	user, err := models.GetUserById(uint(userId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to get user",
		})
	}

	dtoUser := dtos.ParseUserToResponseDto(user)
	return c.JSON(fiber.Map{
		"data": &dtoUser,
	})
}
