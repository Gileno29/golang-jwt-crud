package model

import (
	"fmt"

	"github.com/Gileno29/golang-jwt-crud/src/configuration/logger"
	"github.com/Gileno29/golang-jwt-crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("init CreateUser model", zap.String("journey", "createuser"))
	ud.EncryptPassword()

	fmt.Println(ud)
	return nil
}
