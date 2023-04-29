package controllers

import (
	"log"
	"strconv"

	"github.com/fahad-md-kamal/fiber-blogs/users/dtos"
	"github.com/fahad-md-kamal/fiber-blogs/users/models"
	"github.com/fahad-md-kamal/fiber-blogs/utils"
	"github.com/gofiber/fiber/v2"
)

func AddUserHandler(c *fiber.Ctx) error {

	var userCreateDto dtos.UserCreateDto

	if err := c.BodyParser(&userCreateDto); err != nil {
		log.Printf("Failed to perse provided | Data: %s | Error: %s", string(c.Body()), err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid data to create user",
		})
	}

	if errors, ok := userCreateDto.ValidateUserCreateDto(); !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors})
	}

	UserToCreate := userCreateDto.ParseFromDto()
	if _, ok := UserToCreate.GeneratePasswordHash(); !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to generate password hash",
		})
	}

	if message, ok := UserToCreate.ValidateUserExists(); ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": message,
		})
	}

	if err := UserToCreate.Save(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to create user",
		})
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
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Failed to get user's list",
		})
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
		log.Printf("Error parsing userId: %s | Error: %s", c.Params("id"), err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid User Id",
		})
	}

	user, err := models.GetUserById(uint(userId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Failed to get user",
		})
	}

	dtoUser := dtos.ParseUserToResponseDto(user)
	return c.JSON(fiber.Map{
		"data": &dtoUser,
	})
}

func UpdateUserHandler(c *fiber.Ctx) error {
	userId, err := strconv.ParseUint(c.Params("id"), 10, 0)
	if err != nil {
		log.Printf("Error parsing userId: %s | Error: %s", c.Params("id"), err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid User Id",
		})
	}

	var userUpdateDto dtos.UserUpdateDto
	if err := c.BodyParser(&userUpdateDto); err != nil {
		log.Printf("Failed to perse provided | Data: %s | Error: %s", string(c.Body()), err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse provided data",
		})
	}

	if errors, ok := userUpdateDto.ValidateUserUpdateDto(); !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": errors,
		})
	}

	userToUpdate, err := models.GetUserById(uint(userId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Failed to get user",
		})
	}

	userCheckParams := models.UserCheckParams{
		UserId: userToUpdate.ID,
		Email:  userUpdateDto.Email,
	}
	msg, exists := models.ValidateUserExistsWithEmailOrUsername(userCheckParams)

	if exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": msg,
		})
	}

	updatedUser, err := userToUpdate.UpdateUser(&userUpdateDto)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}

	dtoUser := dtos.ParseUserToResponseDto(updatedUser)
	return c.JSON(fiber.Map{
		"data": &dtoUser,
	})
}

func DeleteUserHandler(c *fiber.Ctx) error {
	userId, err := strconv.ParseUint(c.Params("id"), 10, 0)
	if err != nil {
		log.Printf("Error parsing userID: %s | Error: %s", c.Params("id"), err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid User Id",
		})
	}

	userToDelete, err := models.GetUserById(uint(userId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Failed to get user",
		})
	}

	if err := userToDelete.DeleteUser(); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Couldn't delete user",
		})
	}
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}
