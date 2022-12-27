package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/wrapper/kratos/login"
	"github.com/sdslabs/nymeria/pkg/wrapper/kratos/registration"
)

func HandleOIDCLogin(c *gin.Context) {
	log.Logger.Debug("Get Google Login")
	cookie, flowID, csrf_token, err := login.InitializeLoginFlowWrapper()

		if err != nil {
		log.ErrorLogger("Intialize OIDC Login Failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}
	c.SetCookie("googlelogin_flow", cookie, 3600, "/", "localhost", false, true)

}

func HandleOIDCRegister(c *gin.Context) {
		cookie, flowID, csrf_token, err := registration.InitializeRegistrationFlowWrapper()

	if err != nil {
		log.ErrorLogger("Kratos get OIDC registration flow failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.SetCookie("registration_flow", cookie, 3600, "/", "localhost", false, true)

}