package service

import (
	"github.com/Ricardolv/mvc-api/src/config/logger"
	"github.com/Ricardolv/mvc-api/src/config/rest_err"
	"github.com/Ricardolv/mvc-api/src/model"

	"go.uber.org/zap"
)

func (ud *userDomainService) Create(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init createUser model.",
		zap.String("journey", "createUser"))

	user, _ := ud.FindByEmail(userDomain.GetEmail())
	if user != nil {
		return nil, rest_err.NewBadRequestError("Email is already registered in another account")
	}

	userDomain.EncryptPassword()
	userDomainRepository, err := ud.userRepository.Create(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info(
		"CreateUser service executed successfully",
		zap.String("userId", userDomainRepository.GetID()),
		zap.String("journey", "createUser"))
	return userDomainRepository, nil
}
