package server

import (
	"github.com/fahad-md-kamal/fiber-blogs/configs"
	"github.com/fahad-md-kamal/fiber-blogs/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupAndListen() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://gofiber.io, https://gofiber.net, *",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	users.UsersRouts(app)
	app.Listen(configs.ENVs.ServingPort)
}
