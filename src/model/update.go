package model

import (
	"github.com/Ricardolv/mvc-api/src/config/logger"
	"github.com/Ricardolv/mvc-api/src/config/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) Update(string) *rest_err.RestErr {
	logger.Info("Init update model", zap.String("journey", "updateUser"))

	return nil
}
