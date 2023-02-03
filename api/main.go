package api

import (
	"github.com/gin-gonic/gin"
	c "github.com/sdslabs/nymeria/pkg/controller/admin"
	"github.com/sdslabs/nymeria/pkg/middleware"
)

func Start() {
	r := gin.Default()
	// k := m.NewMiddleware()

	// r.Use(k.Session())

	r.GET("/ping", middleware.OnlyAdmin, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/login", HandleGetLoginFlow)
	r.POST("/login", HandlePostLoginFlow)

	r.POST("/create-identity", c.CreateIdentity)
	r.GET("/get-identity", c.GetIdentity)
	r.POST("/delete-identity", c.DeleteIdentity)
	r.GET("/list-identity", c.ListIdentity)
	r.PUT("/update-identity/ban", c.UpdateBanIdentity)

	r.GET("/register", HandleGetRegistrationFlow)
	r.POST("/register", HandlePostRegistrationFlow)

	r.GET("/logout", HandleGetLogoutFlow)
	r.POST("/logout", HandlePostLogoutFlow)

	r.GET("/status", HandleStatus)

	r.GET("/recovery", HandleGetRecoveryFlow)
	r.POST("/recovery", HandlePostRecoveryFlow)

	r.GET("/settings", HandleGetSettingsFlow)
	r.POST("/settings", HandlePostSettingsFlow)

	r.GET("/verification", HandleGetVerificationFlow)
	r.POST("/verification", HandlePostVerificationFlow)

	r.Run(":9898")
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
