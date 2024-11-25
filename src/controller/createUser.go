package controller

import (
	"github.com/Gileno29/golang-jwt-crud/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Rota invocada de forma errada")
	c.JSON(err.Code, err)
}
