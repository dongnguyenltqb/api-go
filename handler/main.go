package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run(port int) error {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/me", Authenicated, Authenicated)

	return r.Run(fmt.Sprintf(":%v", port))
}
