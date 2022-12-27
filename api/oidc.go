package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/wrapper/kratos/login"
	"github.com/sdslabs/nymeria/pkg/wrapper/kratos/registration"
	"github.com/sdslabs/nymeria/pkg/wrapper/kratos/oidc"
)

func HandleOIDCLogin(c *gin.Context) {
	log.Logger.Debug("Get OIDC Login")
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
	log.Logger.Debug("Get OIDC Registration")
	provider := c.Param("provider")
	if provider == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "provider not found",
		})
		return
	}
		cookie, flowID, csrf_token, err := registration.InitializeRegistrationFlowWrapper()

	if err != nil {
		log.ErrorLogger("Kratos get OIDC registration flow failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.SetCookie("OIDC_registration_flow", cookie, 3600, "/", "localhost", false, true)
	//In case we need to separate the flows so setting and getting cookies simultaneously
    afterCookie, err := c.Cookie("OIDC_registration_flow")

	if err != nil {
		log.ErrorLogger("Cookie not found", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "csrf cookie not found",
		})
		return
	}	
		session, err := oidc.SubmitOIDCRegistrationFlowWrapper(provider, afterCookie,  flowID, csrf_token)

	if err != nil {
		log.ErrorLogger("Kratos post registration flow failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}
	c.SetCookie("sdslabs_session", session, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"status": "created",
	})

}