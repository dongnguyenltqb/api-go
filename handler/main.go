package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run(port int) error {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "welcome to my API")
	})
	// group user
	groupUser := r.Group("/api/users")
	{
		groupUser.GET("/me", Authenicated, getMe)
		groupUser.POST("/", createUserHandler)
	}

	return r.Run(fmt.Sprintf(":%v", port))
}
