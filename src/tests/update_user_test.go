package tests

import (
	"context"
	"encoding/json"
	"github.com/Ricardolv/mvc-api/src/controller/request"
	"github.com/Ricardolv/mvc-api/src/model/repository/entity"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestUpdateUser(t *testing.T) {
	recorder := httptest.NewRecorder()
	ctx := GetTestGinContext(recorder)
	id := primitive.NewObjectID()

	_, err := Database.
		Collection("test_user").
		InsertOne(context.Background(), bson.M{
			"_id":   id,
			"name":  "OLD_NAME",
			"age":   10,
			"email": "test@test.com"})
	if err != nil {
		t.Fatal(err)
		return
	}

	param := []gin.Param{
		{
			Key:   "id",
			Value: id.Hex(),
		},
	}

	userRequest := request.UserUpdateRequest{
		Name: "HUNCODING_NEW_NAME",
		Age:  int8(71),
	}

	b, _ := json.Marshal(userRequest)
	stringReader := io.NopCloser(strings.NewReader(string(b)))

	MakeRequest(ctx, param, url.Values{}, "PUT", stringReader)
	UserController.Update(ctx)

	assert.EqualValues(t, http.StatusOK, recorder.Result().StatusCode)

	userEntity := entity.UserEntity{}

	filter := bson.D{{Key: "_id", Value: id}}
	_ = Database.
		Collection("test_user").
		FindOne(context.Background(), filter).Decode(&userEntity)

	assert.EqualValues(t, userRequest.Name, userEntity.Name)
	assert.EqualValues(t, userRequest.Age, userEntity.Age)
}
