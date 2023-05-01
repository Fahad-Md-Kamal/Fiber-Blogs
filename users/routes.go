package users

import (
	"github.com/fahad-md-kamal/fiber-blogs/middlewares"
	"github.com/fahad-md-kamal/fiber-blogs/users/controllers"
	"github.com/gofiber/fiber/v2"
)

func UsersRouts(app *fiber.App) {
	router := app.Group("users", middlewares.JwtMiddleware())

	router.Post("/", controllers.AddUserHandler)
	router.Get("/", controllers.GetUsersListHandler)
	router.Get("/:id", controllers.GetUserDetailHandler)
	router.Put("/:id", controllers.UpdateUserHandler)
	router.Delete("/:id", controllers.DeleteUserHandler)

	unProtectedRoute := app.Group("")
	unProtectedRoute.Get("/logout", middlewares.JwtMiddleware(), controllers.LogoutHandler)
	unProtectedRoute.Post("/login", controllers.LoginHandler)
}
