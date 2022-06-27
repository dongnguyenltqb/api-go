package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Authenicated(c *gin.Context) {
	if c.Keys["checked"] == nil {
		c.Set("checked", 0)
	}
	c.Set("checked", c.Keys["checked"].(int)+1)

	fmt.Println("authenticated check ", c.Keys["checked"])
	c.Next()
}
