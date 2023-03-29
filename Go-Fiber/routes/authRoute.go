package routes

import (
	"github.com/gofiber/fiber"
	"github.com/shuklaritvik06/GoProjects/fiber/controllers"
)

func AuthRoute(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/refreshtoken", controllers.RefreshToken)
}
