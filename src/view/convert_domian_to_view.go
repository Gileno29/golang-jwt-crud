package view

import (
	"github.com/Gileno29/golang-jwt-crud/src/controller/model/response"
	"github.com/Gileno29/golang-jwt-crud/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
