package service

import (
	"fmt"

	"github.com/Gileno29/golang-jwt-crud/src/configuration/logger"
	"github.com/Gileno29/golang-jwt-crud/src/configuration/rest_err"
	"github.com/Gileno29/golang-jwt-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("init CreateUser model", zap.String("journey", "createuser"))
	userDomain.EncryptPassword()

	fmt.Println(ud)
	return nil
}
