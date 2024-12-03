package controller

import (
	"net/http"

	"github.com/Gileno29/golang-jwt-crud/src/configuration/logger"
	"github.com/Gileno29/golang-jwt-crud/src/configuration/validation"
	"github.com/Gileno29/golang-jwt-crud/src/controller/model/request"
	"github.com/Gileno29/golang-jwt-crud/src/model"
	"github.com/Gileno29/golang-jwt-crud/src/model/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	service := service.NewUserDomainService()

	if err := service.CreateUser(domain); err != nil {
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

	logger.Info("User create seccessfully", zap.String("journey", "CreateUser"))

	c.String(http.StatusOK, "")
}
