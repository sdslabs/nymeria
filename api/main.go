package api

import (
	"github.com/gin-gonic/gin"
	c "github.com/sdslabs/nymeria/pkg/controller/admin"
	m "github.com/sdslabs/nymeria/pkg/middleware"
)

func Start() {
	r := gin.Default()
	k := m.NewMiddleware()

	r.Use(k.Session())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/create-identity", c.CreateIdentity)
	r.GET("/get-identity", c.GetIdentity)
	r.POST("/delete-identity", c.DeleteIdentity)

	r.GET("/register", HandleGetRegistrationFlow)
	r.POST("/register", HandlePostRegistrationFlow)

	r.Run()
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
