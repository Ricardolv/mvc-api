package service

import (
	"github.com/Ricardolv/mvc-api/src/config/logger"
	"github.com/Ricardolv/mvc-api/src/config/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteService(
	userId string) *rest_err.RestErr {

	logger.Info("Init deleteUser model.", zap.String("journey", "deleteUser"))

	err := ud.userRepository.Delete(userId)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "deleteUser"))
		return err
	}

	logger.Info("deleteUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"))
	return nil
}
