package service

import (
	"github.com/Ricardolv/mvc-api/src/config/logger"
	"github.com/Ricardolv/mvc-api/src/config/rest_err"
	"github.com/Ricardolv/mvc-api/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindBy(string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findBy model", zap.String("journey", "findByUser"))

	return nil, nil
}
