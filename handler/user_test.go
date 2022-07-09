package handler

import (
	"context"
	"learn/common"
	"learn/config"
	"learn/entity"
	"learn/util"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/go-resty/resty/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

const PORT = 1234
const BASE_URL = "http://localhost:1234"

func TestUser(t *testing.T) {
	config.Load()
	svc := Run(PORT)
	defer svc.Shutdown(context.Background())

	g := Goblin(t)
	g.Describe("user api", func() {
		g.Before(func() {
			common.GetDB().Create(&entity.User{
				Email: "nhnguhoc@gmail.com",
				Name:  util.Ref("dongnguyenltqb"),
				Age:   25,
			})
		})
		g.After(func() {
			common.GetDB().Where(map[string]any{"email": "nhnguhoc@gmail.com"}).Delete(&entity.User{})
		})
		g.It("create user", func() {
			createObject := createUser{
				Email: "nhnguhoc@gmail.com",
				Age:   18,
			}
			response := new(APIResponse)
			user := new(entity.User)
			resty.New().R().SetBody(createObject).SetResult(response).Post(BASE_URL + "/api/users")
			mapstructure.Decode(response.Data, user)
			g.Assert(user.Email == createObject.Email)
			g.Assert(user.Age == createObject.Age)
			g.Assert(user.Name == nil)
		})

		g.It("get me", func() {
			response := new(APIResponse)
			res, _ := resty.New().R().SetResult(response).SetError(response).Get(BASE_URL + "/api/users/me")
			assert.True(t, res.StatusCode() == http.StatusUnauthorized, "status must be 401")
			assert.True(t, response.Success == false, "response must failed")
			assert.True(t, response.Message == "Unauthenticated.", "message must be Unauthenticated.")
		})
	})
}
