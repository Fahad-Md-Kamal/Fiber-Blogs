package controllers

import "github.com/gofiber/fiber/v2"

func GetUsersListHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Yeee!!!, Fiber Project has started",
	})
}
