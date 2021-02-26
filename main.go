package main

import (
	"github.com/arfan21/gofiber-tes/config"
	"github.com/arfan21/gofiber-tes/controllers"
	"github.com/arfan21/gofiber-tes/helpers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := config.ConnectDB()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(helpers.Response(c, fiber.StatusOK, "success", "hello world tes", nil))
	})

	userCtrl := controllers.NewUserController(db)
	userCtrl.Routes(app)

	app.Listen(":8000")
}
