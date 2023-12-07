package model

import (
	"github.com/Ricardolv/mvc-api/src/config/logger"
	"github.com/Ricardolv/mvc-api/src/config/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) Delete(string) *rest_err.RestErr {
	logger.Info("Init delete model", zap.String("journey", "deleteUser"))

	return nil
}
