package handler

import "github.com/gin-gonic/gin"

func UnAuthenticated(c *gin.Context) {
	c.AbortWithStatusJSON(401, APIResponse{
		Success: false,
		Message: "Unauthenticated.",
	})
}

func BadRequest(c *gin.Context, err error) {
	c.AbortWithStatusJSON(400, APIResponse{
		Success: false,
		Message: err.Error(),
	})
}

func ResponseOK(c *gin.Context, response interface{}) {
	c.JSON(200, APIResponse{
		Success: true,
		Data:    response,
	})
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

func GET(path string, group *gin.RouterGroup, funcs ...gin.HandlerFunc) {
	group.GET(path, funcs...)
}

func POST(path string, group *gin.RouterGroup, funcs ...gin.HandlerFunc) {
	group.POST(path, funcs...)
}

func PUT(path string, group *gin.RouterGroup, funcs ...gin.HandlerFunc) {
	group.PUT(path, funcs...)
}

func DELETE(path string, group *gin.RouterGroup, funcs ...gin.HandlerFunc) {
	group.DELETE(path, funcs...)
}
