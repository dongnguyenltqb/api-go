package handler

import (
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
	go Run(PORT)
	createObject := createUser{
		Email: "nhnguhoc@gmail.com",
		Age:   18,
	}
	response := new(APIResponse)
	user := new(entity.User)

	resty.New().R().SetBody(createObject).SetResult(response).Post(BASE_URL + "/api/users")
	mapstructure.Decode(response.Data, user)
	t.Logf("name: %+v", user)
	assert.True(t, user.Email == createObject.Email)
	assert.True(t, user.Age == createObject.Age)
	assert.True(t, user.Name == nil, "name must be NIL")
}
