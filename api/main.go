package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/register", HandleGetRegistrationFlow)

	if err := r.Run(); err != nil {
		fmt.Println(err)
	}
}
