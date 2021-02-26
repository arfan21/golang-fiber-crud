package repository

import (
	"context"
	"github.com/arfan21/gofiber-tes/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id string) (*models.User, error)
	GetAll() (*[]models.User, error)
	Update(id string, user *models.User) error
	Delete(id string) error
}

type userRepo struct {
	ctx context.Context
	db  *mongo.Database
}

func NewUserRepo(ctx context.Context, db *mongo.Database) UserRepository {
	repo := &userRepo{ctx: ctx, db: db}

	return repo
}

func (repo *userRepo) Create(user *models.User) error {
	user.ID = primitive.NewObjectID()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	_, err := repo.db.Collection("user").InsertOne(repo.ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func (repo *userRepo) GetByID(id string) (*models.User, error) {
	user := new(models.User)
	ObjectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	err = repo.db.Collection("user").FindOne(repo.ctx, bson.M{"_id": ObjectID}).Decode(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *userRepo) GetAll() (*[]models.User, error) {
	result, err := repo.db.Collection("user").Find(repo.ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer result.Close(repo.ctx)

	users := make([]models.User, 0)

	for result.Next(repo.ctx) {
		row := new(models.User)

		err := result.Decode(&row)

		if err != nil {
			return nil, err
		}

		row.Password = ""

		users = append(users, *row)
	}

	return &users, nil
}

func (repo *userRepo) Update(id string, user *models.User) error {
	ObjectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	err = repo.db.Collection("user").FindOneAndUpdate(repo.ctx, bson.M{"_id": ObjectID}, bson.M{"$set": user}).Err()

	if err != nil {
		return err
	}

	user.ID = ObjectID

	return nil
}

func (repo *userRepo) Delete(id string) error {
	ObjectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	err = repo.db.Collection("user").FindOneAndDelete(repo.ctx, bson.M{"_id": ObjectID}).Err()

	if err != nil {
		return err
	}

	return nil
}
