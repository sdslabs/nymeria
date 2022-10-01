package api

import (
    "github.com/gin-gonic/gin"
	m"github.com/sdslabs/nymeria/pkg/middleware"
	c"github.com/sdslabs/nymeria/pkg/controller/admin"
<<<<<<< HEAD
=======
    l"github.com/sdslabs/nymeria/pkg/controller/login"
>>>>>>> bea42114aa1a480459ebcccc6109d01ffca3d767
)


func Start() {

<<<<<<< HEAD
<<<<<<< HEAD
	r.GET("/register", controller.Registration)
=======
=======
>>>>>>> bea42114aa1a480459ebcccc6109d01ffca3d767
    r := gin.Default()
    k := m.NewMiddleware()

    r.Use(k.Session())

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

<<<<<<< HEAD
	//r.POST("/create-identity", c.CreateIdentity)
    r.POST("/create-identity", c.CreateIdentity)
	r.GET("/get-identity", c.GetIdentity)
	r.POST("/delete-identity", c.DeleteIdentity)

	
	r.Run()
<<<<<<< HEAD
>>>>>>> e06a244 (Added identities controllers)

=======
>>>>>>> a204a8e (implemented delete identities controller)
=======
    r.POST("/create-identity", c.CreateIdentity)
	r.GET("/get-identity", c.GetIdentity)
	r.POST("/delete-identity", c.DeleteIdentity)
    r.GET("/login", l.InitializeSelfServiceLoginFlow)

	
	r.Run()
>>>>>>> bea42114aa1a480459ebcccc6109d01ffca3d767
    // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}