package users

import (
	"github.com/fahad-md-kamal/fiber-blogs/users/controllers"
	"github.com/gofiber/fiber/v2"
)

func UsersRouts(app *fiber.App) {
	router := app.Group("users")

	router.Post("/", controllers.AddUserHandler)
	router.Get("/", controllers.GetUsersListHandler)
	router.Get("/:id", controllers.GetUserDetailHandler)
	router.Put("/:id", controllers.UpdateUserHandler)
	router.Delete("/:id", controllers.DeleteUserHandler)

	unProtectedRoute := app.Group("")
	unProtectedRoute.Post("/login", controllers.LoginHandler)
	unProtectedRoute.Get("/logout", controllers.LogoutHandler)
}
