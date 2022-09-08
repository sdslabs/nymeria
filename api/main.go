package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sdslabs/nymeria/pkg/controller"
)

func Start() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/register", controller.Registration)

	if err := r.Run(); err != nil {
		fmt.Println(err)
	}
}
