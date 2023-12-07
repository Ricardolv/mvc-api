package repository

import (
	"github.com/Ricardolv/mvc-api/src/config/rest_err"
	"github.com/Ricardolv/mvc-api/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	Create(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)

	Update(
		userId string,
		userDomain model.UserDomainInterface,
	) *rest_err.RestErr

	Delete(
		userId string,
	) *rest_err.RestErr

	FindByEmail(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindByEmailAndPassword(
		email string,
		password string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindByID(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)
}
