package main

import (
	"context"
	"log"

	"github.com/Ricardolv/mvc-api/src/config/database/mongodb"
	"github.com/Ricardolv/mvc-api/src/config/logger"
	"github.com/Ricardolv/mvc-api/src/controller"
	"github.com/Ricardolv/mvc-api/src/controller/routes"
	"github.com/Ricardolv/mvc-api/src/model/repository"
	"github.com/Ricardolv/mvc-api/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/Ricardolv/mvc-api/docs"
)

// @title MVC Api em Go | Richard
// @version 1.0
// @description API for crud operations on users
// @host localhost:8080
// @BasePath /users
// @schemes http
// @license MIT
func main() {
	logger.Info("About to start user application")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	controller := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, controller)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
