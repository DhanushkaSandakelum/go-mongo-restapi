package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-mongo-restapi/controllers"
)

func UserRoute(app *fiber.App){
	app.Post("/user", controllers.CreateUser)
}