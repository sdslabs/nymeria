package api

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/sdslabs/nymeria/pkg/middleware"
)

func Start() {
	r := gin.Default()
	// Set up CORS middleware
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(config))

	// r.Use(k.Session())

	r.GET("/ping", middleware.OnlyAdmin, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/login", HandleGetLoginFlow)
	r.POST("/login", HandlePostLoginFlow)
	r.GET("/mfa", HandleGetMFAFlow)
	r.POST("/mfa", HandlePostMFAFlow)

	r.POST("/create-identity", middleware.OnlyAdmin, HandleCreateIdentityFlow)
	r.GET("/get-identity", middleware.OnlyAdmin, HandleGetIdentityFlow)
	r.POST("/delete-identity", middleware.OnlyAdmin, HandleDeleteIdentityFlow)
	r.GET("/list-identity", middleware.OnlyAdmin, HandleListIdentity)
	r.PUT("/update-identity/ban", middleware.OnlyAdmin, HandleBanIdentity)
	r.PUT("/update-identity/remove-ban", middleware.OnlyAdmin, HandleRemoveBanIdentity)
	r.PUT("/update-identity/switch-roles", middleware.OnlyAdmin, HandleRoleSwitch)

	r.GET("/register", HandleGetRegistrationFlow)
	r.POST("/register", HandlePostRegistrationFlow)

	r.GET("/logout", HandleGetLogoutFlow)
	r.POST("/logout", HandlePostLogoutFlow)

	r.GET("/status", HandleStatus)

	r.GET("/recovery", HandleGetRecoveryFlow)
	r.POST("/recovery", HandlePostRecoveryFlow)

	r.GET("/settings", HandleGetSettingsFlow)
	r.POST("/updateprofile", HandleUpdateProfile)
	r.POST("/changepassword", HandleChangePassword)
	r.POST("/toggletotp", HandleToggleTOTP)

	r.GET("/verification", HandleGetVerificationFlow)
	r.POST("/verification", HandlePostVerificationFlow)

	r.POST("/get_profile", HandlePostProfile)
	r.POST("/verify_app", middleware.HandleAppAuthorization, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Authorized",
		})
	})
	r.GET("/application", HandleGetApplication)
	r.POST("/application", HandlePostApplication)
	r.PUT("/application", HandlePutApplication)
	r.DELETE("/application", HandleDeleteApplication)

	r.POST("/update-client-secret", HandleUpdateClientSecret)
	r.Run(":9898")
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
