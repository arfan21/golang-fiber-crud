package user

import (
	"github.com/arfan21/golang-fiber-crud/models"
	"github.com/arfan21/golang-fiber-crud/repository/mongo/user"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(user *models.User) error
	GetByID(id string) (*models.User, error)
	GetAll() (*[]models.User, error)
	Update(id string, user *models.User) error
	Delete(id string) error
}

type userService struct {
	repo user.UserRepository
}

func NewUserService(repo user.UserRepository) UserService {
	return &userService{repo}
}

func (service *userService) Create(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return service.repo.Create(user)
}

func (service *userService) GetByID(id string) (*models.User, error) {
	return service.repo.GetByID(id)
}


func (service *userService) GetAll() (*[]models.User, error){
	return service.repo.GetAll()
}

func (service *userService) Update(id string, user *models.User) error{
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return service.repo.Update(id, user)
}

func (service *userService) Delete(id string) error {
	return service.repo.Delete(id)
}