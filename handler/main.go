package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createUser struct {
	Name string `json:"name,omitempty"`
	Age  *int   `json:"age,omitempty"`
}

func createUserHandler(c *gin.Context) {
	p := new(createUser)

	if err := c.BindJSON(p); err != nil {
		panic(err)
	}
	fmt.Printf("%+v", p)
	// fmt.Printf("age = %d  , name = %s", *p.Age, p.Name)
	c.JSON(200, p)
}

func Run(port int) error {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/me", Authenicated, Authenicated, createUserHandler)

	return r.Run(fmt.Sprintf(":%v", port))
}
