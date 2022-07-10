package handler

import (
	"fmt"
	"learn/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run(port int) *http.Server {

	gin.SetMode(config.Get().GinMode)
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
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}
	fmt.Printf("server is listening on port %d\n", port)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	return srv
}
