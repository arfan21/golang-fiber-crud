package main

import (
	"github.com/arfan21/golang-fiber-crud/config"
	"github.com/arfan21/golang-fiber-crud/controllers/http/user"
	"github.com/arfan21/golang-fiber-crud/helpers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := config.ConnectDB()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(helpers.Response(c, fiber.StatusOK, "success", "hello world tes", nil))
	})

	userCtrl := user.NewUserController(db)

	userCtrl.Routes(app)

	app.Listen(":8000")
}
