package helpers

import "github.com/gofiber/fiber/v2"

type res struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Response(c *fiber.Ctx, code int, status string, message string, data interface{}) error {
	return c.Status(code).JSON(res{status, message, data})
}
