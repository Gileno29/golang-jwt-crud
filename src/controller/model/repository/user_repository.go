package repository

import (
	"github.com/Gileno29/golang-jwt-crud/src/configuration/rest_err"
	"github.com/Gileno29/golang-jwt-crud/src/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		databaseConection: database,
	}
}

type userRepository struct {
	databaseConection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
}
