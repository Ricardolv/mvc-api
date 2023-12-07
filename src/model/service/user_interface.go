package service

import (
	"github.com/Ricardolv/mvc-api/src/config/rest_err"
	"github.com/Ricardolv/mvc-api/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	Create(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	Update(string, model.UserDomainInterface) *rest_err.RestErr
	FindBy(id string) (model.UserDomainInterface, *rest_err.RestErr)
	Delete(string) *rest_err.RestErr
}
