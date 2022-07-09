package handler

import (
	"learn/entity"

	"github.com/gin-gonic/gin"
)

type createUser struct {
	Name  *string `json:"name,omitempty" binding:"omitempty,min=1"`
	Email string  `json:"email,omitempty" binding:"required,email"`
	Age   int     `json:"age,omitempty" binding:"required,min=1"`
}

func createUserHandler(c *gin.Context) {
	p := new(createUser)
	if err := c.ShouldBindJSON(p); err != nil {
		BadRequest(c, err)
		return
	}
	user := entity.CreateUser(&entity.User{
		Email: p.Email,
		Name:  p.Name,
		Age:   p.Age,
	})
	ResponseOK(c, user)
}

func getMe(c *gin.Context) {
	ResponseOK(c, c.Keys["user_id"])
}
