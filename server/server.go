package server

import (
	"github.com/fahad-md-kamal/fiber-blogs/configs"
	"github.com/fahad-md-kamal/fiber-blogs/users"
	"github.com/gofiber/fiber/v2"
)

func SetupAndListen() {

	app := fiber.New()

	users.UsersRouts(app)
	app.Listen(configs.ENVs.ServingPort)
}
