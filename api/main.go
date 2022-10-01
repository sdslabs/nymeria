package api

import (
    "github.com/gin-gonic/gin"
	m"github.com/sdslabs/nymeria/pkg/middleware"
	c"github.com/sdslabs/nymeria/pkg/controller/admin"
    l"github.com/sdslabs/nymeria/pkg/controller/login"
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
    r.GET("/login", l.InitializeSelfServiceLoginFlow)

	
	r.Run()
    // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}