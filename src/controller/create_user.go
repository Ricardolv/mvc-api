package controller

import (
	"net/http"

	"github.com/Ricardolv/mvc-api/src/config/logger"
	"github.com/Ricardolv/mvc-api/src/config/validation"
	"github.com/Ricardolv/mvc-api/src/controller/request"
	"github.com/Ricardolv/mvc-api/src/model"
	"github.com/Ricardolv/mvc-api/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomanInterface model.UserDomainInterface
)

func (uc *userControllerInterface) Create(c *gin.Context) {

	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	domainResult, err := uc.service.Create(domain)
	if err != nil {
		logger.Error(
			"Error trying to call CreateUser service",
			err,
			zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"CreateUser controller executed successfully",
		zap.String("userId", ""),
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
