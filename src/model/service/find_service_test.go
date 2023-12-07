package service

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/Ricardolv/mvc-api/src/config/rest_err"
	"github.com/Ricardolv/mvc-api/src/model"
	"github.com/Ricardolv/mvc-api/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_FindUserByIDServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_an_user_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 50)
		userDomain.SetID(id)

		repository.EXPECT().FindByID(id).Return(userDomain, nil)

		userDomainReturn, err := service.FindByIDService(id)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetID(), id)
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainReturn.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		repository.EXPECT().FindByID(id).Return(nil, rest_err.NewNotFoundError("user not found"))
		userDomainReturn, err := service.FindByIDService(id)

		assert.Nil(t, userDomainReturn)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "user not found")
	})
}

func TestUserDomainService_FindUserByEmailServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_an_user_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "test@success.com"

		userDomain := model.NewUserDomain(email, "test", "test", 50)
		userDomain.SetID(id)

		repository.EXPECT().FindByEmail(email).Return(userDomain, nil)

		userDomainReturn, err := service.FindByEmailService(email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetID(), id)
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainReturn.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(t *testing.T) {
		email := "test@error.com"

		repository.EXPECT().FindByEmail(email).Return(nil, rest_err.NewNotFoundError("user not found"))
		userDomainReturn, err := service.FindByEmailService(email)

		assert.Nil(t, userDomainReturn)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "user not found")
	})
}

func TestUserDomainService_FindUserByEmailAndPasswordServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := &userDomainService{repository}

	t.Run("when_exists_an_user_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "test@success.com"
		password := strconv.FormatInt(rand.Int63(), 10)

		userDomain := model.NewUserDomain(email, password, "test", 50)
		userDomain.SetID(id)

		repository.EXPECT().FindByEmailAndPassword(email, password).Return(userDomain, nil)

		userDomainReturn, err := service.findByEmailAndPasswordService(email, password)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetID(), id)
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainReturn.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(t *testing.T) {
		email := "test@error.com"
		password := strconv.FormatInt(rand.Int63(), 10)

		repository.EXPECT().FindByEmailAndPassword(email, password).Return(nil, rest_err.NewNotFoundError("user not found"))
		userDomainReturn, err := service.findByEmailAndPasswordService(email, password)

		assert.Nil(t, userDomainReturn)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "user not found")
	})
}
