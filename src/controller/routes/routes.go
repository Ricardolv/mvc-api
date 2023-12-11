package routes

import (
	"github.com/Ricardolv/mvc-api/src/controller"
	"github.com/Ricardolv/mvc-api/src/model"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(
	r *gin.RouterGroup,
	controller controller.UserControllerInterface,
) {

	r.GET("/users/:id", model.VerifyTokenMiddleware, controller.FindByID)
	r.GET("/users/email/:email", model.VerifyTokenMiddleware, controller.FindByEmail)
	r.POST("/users", model.VerifyTokenMiddleware, controller.Create)
	r.PUT("/users/:id", model.VerifyTokenMiddleware, controller.Update)
	r.DELETE("/users/:id", model.VerifyTokenMiddleware, controller.Delete)

	r.POST("/login", controller.LoginUser)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
