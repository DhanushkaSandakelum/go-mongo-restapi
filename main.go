package main

import(
	"go-mongo-restapi/configs"
	"go-mongo-restapi/routes"
	"github.com/gofiber/fiber/v2"
)

func main()  {
	app := fiber.New()

	// Run DB
	configs.ConnectDB()

	// Routes
	routes.UserRoute(app)

	app.Listen(":3333")
}