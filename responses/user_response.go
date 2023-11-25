package responses

import "github.com/gofiber/fiber/v2"

type UserResponse struct {
	Status  int
	Message string
	Data    *fiber.Map
}
