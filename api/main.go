package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sdslabs/nymeria/pkg/middleware"
)

func Start() {
	r := gin.Default()
	// Set up CORS middleware
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // TODO: Change to production domain
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(config))

	// r.Use(k.Session())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/login", HandleGetLoginFlow)
	r.POST("/login", HandlePostLoginFlow)
	r.GET("/mfa", HandleGetMFAFlow)
	r.POST("/mfa", HandlePostMFAFlow)

	r.GET("/register", HandleGetRegistrationFlow)
	r.POST("/register", HandlePostRegistrationFlow)

	r.GET("/logout", HandleGetLogoutFlow)
	r.POST("/logout", HandlePostLogoutFlow)

	r.GET("/recovery", HandleGetRecoveryFlow)
	r.POST("/recovery", HandlePostRecoveryFlow)
	r.POST("/recovery-code", HandlePostRecoveryCodeFlow)

	r.GET("/settings", HandleGetSettingsFlow)
	r.PATCH("/update-profile", HandleUpdateProfile)
	r.POST("/change-password", HandleChangePassword)
	r.POST("/toggle-totp", HandleToggleTOTP)

	r.GET("/verification", HandleGetVerificationFlow)
	r.POST("/verification", HandlePostVerificationFlow)
	r.POST("/verification-code", HandlePostVerificationCodeFlow)

	r.GET("/get-profile", HandlePostProfile)

	// Verify User Session
	r.GET("/verify-session", HandleVerifySession)

	// Application Authorization
	r.POST("/verify-app", HandleAppAuthorization)

	// Admin Routes
	r.Use(middleware.OnlyAdmin)

	// Identity Management
	r.POST("/create-identity", HandleCreateIdentityFlow)
	r.GET("/get-identity", HandleGetIdentityFlow)
	r.POST("/delete-identity", HandleDeleteIdentityFlow)
	r.GET("/list-identity", HandleListIdentity)
	r.PUT("/update-identity/ban", HandleBanIdentity)
	r.PUT("/update-identity/remove-ban", HandleRemoveBanIdentity)
	r.PUT("/update-identity/switch-roles", HandleRoleSwitch)

	// Application Management
	r.GET("/application", HandleGetApplication)
	r.POST("/application", HandleCreateApplication)
	r.PUT("/application", HandleUpdateApplication)
	r.DELETE("/application", HandleDeleteApplication)
	r.PATCH("/update-client-secret", HandleUpdateClientSecret)
	r.PATCH("/update-client-key", HandleUpdateClientKey)

	r.Run(":9898")
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
