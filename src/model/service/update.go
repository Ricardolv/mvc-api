package service

import (
	"github.com/Ricardolv/mvc-api/src/config/logger"
	"github.com/Ricardolv/mvc-api/src/config/rest_err"
	"github.com/Ricardolv/mvc-api/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) Update(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init update model", zap.String("journey", "updateUser"))

	return nil
}
