package repository

import (
	"context"
	"os"

	"github.com/Gileno29/golang-jwt-crud/src/configuration/logger"
	"github.com/Gileno29/golang-jwt-crud/src/configuration/rest_err"
	"github.com/Gileno29/golang-jwt-crud/src/model"
)

const (
	MONGO_USER_DB = "MONGODB_USER_DB"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init createUser repository")
	collection_name := os.Getenv(MONGO_USER_DB)

	collection := ur.databaseConection.Collection(collection_name)
	value, err := userDomain.GetJSONValue()
	if err != nil {
		return nil, rest_err.NewForbiddenError(err.Error())
	}
	result, err := collection.InsertOne(context.Background(), value)

	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	userDomain.SetID(result.InsetedID.(string))

	return userDomain, nil

}
