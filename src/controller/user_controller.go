package controller

import (
	"github.com/Ricardolv/mvc-api/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(
	serviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	FindByID(c *gin.Context)
	FindByEmail(c *gin.Context)

	Delete(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	LoginUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
