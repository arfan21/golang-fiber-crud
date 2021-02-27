package user

import (
	"context"
	userRepo "github.com/arfan21/golang-fiber-crud/repository/mongo/user"
	userService "github.com/arfan21/golang-fiber-crud/services/user"
	"strings"

	"github.com/arfan21/golang-fiber-crud/helpers"
	"github.com/arfan21/golang-fiber-crud/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController interface {
	Routes(app *fiber.App)
}

type userController struct {
	service userService.UserService
}

func NewUserController(db *mongo.Database) UserController {
	repo := userRepo.NewUserRepo(context.Background(), db)
	service := userService.NewUserService(repo)

	return &userController{service}
}

func (ctrl *userController) Routes(app *fiber.App) {
	app.Post("/user", ctrl.Create)
	app.Get("/user", ctrl.GetAll)
	app.Get("/user/:id", ctrl.GetByID)
	app.Put("/user/:id", ctrl.Update)
	app.Delete("/user/:id", ctrl.Delete)
}

func (ctrl *userController) Create(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return helpers.Response(c, fiber.StatusBadRequest, "error", err.Error(), nil)
	}

	err := ctrl.service.Create(user)

	if err != nil {
		return helpers.Response(c, fiber.StatusInternalServerError, "error", err.Error(), nil)
	}

	return helpers.Response(c, fiber.StatusOK, "success", "sukses membuat akun", user)
}

func (ctrl *userController) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := ctrl.service.GetByID(id)

	if err != nil {
		if strings.Contains(err.Error(), "no documents in result") {
			return helpers.Response(c, fiber.StatusNotFound, "error", err.Error(), nil)
		}
		return helpers.Response(c, fiber.StatusInternalServerError, "error", err.Error(), nil)
	}

	return helpers.Response(c, fiber.StatusOK, "success", "sukses mendapatkan data user", user)
}

func (ctrl *userController) GetAll(c *fiber.Ctx) error {
	users, err := ctrl.service.GetAll()

	if err != nil {
		return helpers.Response(c, fiber.StatusBadRequest, "error", err.Error(), nil)
	}

	return helpers.Response(c, fiber.StatusOK, "success", "sukses mendapatkan semua data user", users)
}

func (ctrl *userController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return helpers.Response(c, fiber.StatusBadRequest, "error", err.Error(), nil)
	}

	err := ctrl.service.Update(id, user)

	if err != nil {
		return helpers.Response(c, fiber.StatusBadRequest, "error", err.Error(), nil)
	}

	return helpers.Response(c, fiber.StatusOK, "success", "sukses memperbarui data user", user)
}

func (ctrl *userController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := ctrl.service.Delete(id)

	if err != nil {
		return helpers.Response(c, fiber.StatusBadRequest, "error", err.Error(), nil)
	}

	return helpers.Response(c, fiber.StatusOK, "success", "sukses menghapus data akun", nil)
}
