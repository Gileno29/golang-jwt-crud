package main

import (
	"context"
	"log"

	"github.com/Gileno29/golang-jwt-crud/src/configuration/database/mongodb"
	"github.com/Gileno29/golang-jwt-crud/src/configuration/logger"
	"github.com/Gileno29/golang-jwt-crud/src/controller"
	"github.com/Gileno29/golang-jwt-crud/src/controller/model/repository"
	"github.com/Gileno29/golang-jwt-crud/src/controller/routes"
	"github.com/Gileno29/golang-jwt-crud/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Esta começando")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Init dependencies
	database, err := mongodb.NewMongoDBCOnection(context.Background())

	if err != nil {
		log.Fatalf("Erro tryng to conect to database, erro=%s", err.Error())
		return
	}
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
