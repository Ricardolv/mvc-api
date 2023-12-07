package view

import (
	"github.com/Ricardolv/mvc-api/src/controller/response"
	"github.com/Ricardolv/mvc-api/src/model"
)

func ConverterDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
