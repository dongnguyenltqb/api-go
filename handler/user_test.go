package handler

import (
	"context"
	"learn/config"
	"learn/entity"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

const PORT = 1234
const BASE_URL = "http://localhost:1234"

func TestCreateUser(t *testing.T) {
	config.Load()
	svc := Run(PORT)
	defer svc.Shutdown(context.Background())

	createObject := createUser{
		Email: "nhnguhoc@gmail.com",
		Age:   18,
	}
	response := new(APIResponse)
	user := new(entity.User)
	resty.New().R().SetBody(createObject).SetResult(response).Post(BASE_URL + "/api/users")
	mapstructure.Decode(response.Data, user)
	assert.True(t, user.Email == createObject.Email)
	assert.True(t, user.Age == createObject.Age)
	assert.True(t, user.Name == nil, "name must be NIL")
}

func TestGetMe(t *testing.T) {
	config.Load()
	svc := Run(PORT)
	defer svc.Shutdown(context.Background())
	response := new(APIResponse)
	resty.New().R().SetResult(response).SetError(response).Get(BASE_URL + "/api/users/me")
	assert.True(t, response.Success == false)
	assert.True(t, response.Message == "Unauthenticated.")

}
