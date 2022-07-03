package handler

import (
	"fmt"
	"learn/util"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authenicated(c *gin.Context) {
	tokens := c.Request.Header["Authorization"]
	fmt.Println("token=", tokens)
	if len(tokens) != 1 {
		UnAuthenticated(c)
		return
	}
	claims := make(jwt.MapClaims)
	if err := util.JwtParse(&claims, tokens[0], []byte("asecretkeyhaha")); err != nil {
		UnAuthenticated(c)
		return
	}
	fmt.Println("claim = ", claims)
	c.Set("user_id", claims["user_id"])
	c.Next()
}
