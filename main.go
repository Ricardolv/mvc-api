package main

import (
	"log"

	"github.com/Ricardolv/mvc-api/src/config/logger"
	"github.com/Ricardolv/mvc-api/src/controller/routes"
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

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
