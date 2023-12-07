package service

import (
	"github.com/Ricardolv/mvc-api/src/config/rest_err"
	"github.com/Ricardolv/mvc-api/src/model"
	"github.com/Ricardolv/mvc-api/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	Create(
		model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)

	Update(
		string, model.UserDomainInterface,
	) *rest_err.RestErr

	FindByID(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindByEmail(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	Delete(string) *rest_err.RestErr
}
