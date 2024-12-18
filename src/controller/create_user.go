package controller

import (
	"net/http"

	"github.com/Gileno29/golang-jwt-crud/src/configuration/logger"
	"github.com/Gileno29/golang-jwt-crud/src/configuration/validation"
	"github.com/Gileno29/golang-jwt-crud/src/controller/model/request"
	"github.com/Gileno29/golang-jwt-crud/src/model"
	"github.com/Gileno29/golang-jwt-crud/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	domainResult, err := uc.service.CreateUser(domain)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Init CreateUser controller", zap.String("journey", "CreateUser"))

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err)
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	logger.Info("User create seccessfully ", zap.Any("Usuario:", domain.GetName()), zap.String("journey", "CreateUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
