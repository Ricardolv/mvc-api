package model

import (
	"github.com/Ricardolv/mvc-api/src/config/logger"
	"github.com/Ricardolv/mvc-api/src/config/rest_err"

	"go.uber.org/zap"
)

func (ud *UserDomain) Create() *rest_err.RestErr {
	logger.Info("Init create model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	logger.Info("End create model", zap.String("journey", "createUser"))
	return nil
}
