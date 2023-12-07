package controller

import (
	"net/http"

	"github.com/Ricardolv/mvc-api/src/config/logger"
	"github.com/Ricardolv/mvc-api/src/config/validation"
	"github.com/Ricardolv/mvc-api/src/controller/request"
	"github.com/Ricardolv/mvc-api/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomanInterface model.UserDomainInterface
)

func Create(c *gin.Context) {

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

	if err := domain.Create(); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info(
		"CreateUser controller executed successfully",
		zap.String("userId", ""),
		zap.String("journey", "createUser"))

	c.String(http.StatusOK, "")
}
