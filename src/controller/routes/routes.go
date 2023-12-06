package routes

import (
	"github.com/Ricardolv/mvc-api/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/users/:id", controller.FindById)
	r.GET("/users/email/:email", controller.FindByEmail)
	r.POST("/users", controller.Create)
	r.PUT("/users/:id", controller.Create)
	r.DELETE("/users/:id", controller.Delete)

}
