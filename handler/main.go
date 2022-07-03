package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UnAuthenticated(c *gin.Context) {
	c.AbortWithStatusJSON(401, gin.H{
		"success": false,
		"message": "Unauthenticated.",
	})
}

func BadRequest(c *gin.Context, err error) {
	c.AbortWithStatusJSON(400, gin.H{
		"success": false,
		"message": err.Error(),
	})
}

func ResponseOK(c *gin.Context, response interface{}) {
	c.JSON(200, gin.H{
		"success": true,
		"data":    response,
	})
}

func Run(port int) error {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "welcome to my API")
	})
	// group user
	groupUser := r.Group("/users")
	{
		groupUser.GET("/me", Authenicated, getMe)
		groupUser.POST("/", createUserHandler)
	}

	return r.Run(fmt.Sprintf(":%v", port))
}
