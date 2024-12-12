package service

import (
	"github.com/Gileno29/golang-jwt-crud/src/configuration/rest_err"
	"github.com/Gileno29/golang-jwt-crud/src/controller/model/repository"
	"github.com/Gileno29/golang-jwt-crud/src/model"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
