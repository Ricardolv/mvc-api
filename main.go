package main

import (
	"log"

	"github.com/Ricardolv/mvc-api/src/config/logger"
	"github.com/Ricardolv/mvc-api/src/controller"
	"github.com/Ricardolv/mvc-api/src/controller/routes"
	"github.com/Ricardolv/mvc-api/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	service := service.NewUserDomainService()
	controller := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, controller)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
