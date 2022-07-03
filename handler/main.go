package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(c *gin.Context, err error) {
	c.Header("Content-Type", "application/json")
	c.JSON(400, gin.H{
		"success": false,
		"message": err.Error(),
	})
}

func ResponseOK(c *gin.Context, response interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(200, response)
}

func Run(port int) error {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "welcome to my API")
	})
	// group user
	groupUser := r.Group("/users")
	{
		groupUser.POST("/me", Authenicated, Authenicated, createUserHandler)
		groupUser.POST("/", createUserHandler)
	}

	return r.Run(fmt.Sprintf(":%v", port))
}
