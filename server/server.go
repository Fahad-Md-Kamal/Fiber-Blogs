package server

import (
	"github.com/fahad-md-kamal/fiber-blogs/users"
	"github.com/gofiber/fiber/v2"
)

func SetupAndListen() {
	app := fiber.New()

	users.UsersRouts(app)
	app.Listen(":3000")
}
