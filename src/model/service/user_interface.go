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
	CreateService(
		model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateService(
		string, model.UserDomainInterface,
	) *rest_err.RestErr

	FindByIDService(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindByEmailService(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	findByEmailAndPasswordService(
		email string,
		password string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	DeleteService(string) *rest_err.RestErr

	LoginUserServices(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, string, *rest_err.RestErr)
}
