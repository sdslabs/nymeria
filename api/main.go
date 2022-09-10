package api

import (

    "github.com/gin-gonic/gin"
	m"github.com/sdslabs/nymeria/pkg/middleware"
	c"github.com/sdslabs/nymeria/pkg/controller"
	//client "github.com/ory/kratos-client-go"
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

	r.GET("/create-identity", c.CreateIdentity)
	r.GET("/get-identity", c.GetIdentity)
	r.GET("/delete-identity", c.DeleteIdentity)

	
	r.Run()

    // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}