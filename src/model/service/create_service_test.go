package service

import (
	"testing"

	"github.com/Ricardolv/mvc-api/src/config/rest_err"
	"github.com/Ricardolv/mvc-api/src/model"
	"github.com/Ricardolv/mvc-api/src/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserDomainService_CreateUserServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_user_already_exists_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("tests@tests.com", "tests", "tests", 50)
		userDomain.SetID(id)

		repository.EXPECT().FindByEmail(userDomain.GetEmail()).Return(userDomain, nil)

		user, err := service.CreateService(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Email is already registered in another account")
	})

	t.Run("when_user_is_not_registered_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("tests@tests.com", "tests", "tests", 50)
		userDomain.SetID(id)

		repository.EXPECT().FindByEmail(userDomain.GetEmail()).Return(nil, nil)

		repository.EXPECT().Create(userDomain).Return(nil,
			rest_err.NewInternalServerError("error trying to create user"))

		user, err := service.CreateService(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to create user")
	})

	t.Run("when_user_is_not_registered_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("tests@tests.com", "tests", "tests", 50)
		userDomain.SetID(id)

		repository.EXPECT().FindByEmail(userDomain.GetEmail()).Return(nil, nil)

		repository.EXPECT().Create(userDomain).Return(userDomain, nil)

		user, err := service.CreateService(userDomain)

		assert.Nil(t, err)
		assert.EqualValues(t, user.GetName(), userDomain.GetName())
		assert.EqualValues(t, user.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, user.GetAge(), userDomain.GetAge())
		assert.EqualValues(t, user.GetID(), userDomain.GetID())
		assert.EqualValues(t, user.GetPassword(), userDomain.GetPassword())
	})
}
