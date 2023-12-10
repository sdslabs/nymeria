package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/sdslabs/nymeria/config"
	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/wrapper/kratos/login"
)

func HandleGetMFAFlow(c *gin.Context) {
	log.Logger.Debug("Get MFA")
	cookie, _ := c.Cookie("sdslabs_session")
	flow_cookie, flowID, csrf_token, err := login.InitializeLoginFlowWrapper("aal2", cookie)

	if err != nil {
		log.ErrorLogger("Initialize MFA Failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.SetCookie("mfa", flow_cookie, 3600, "/", config.NymeriaConfig.URL.Domain, true, true)

	c.JSON(http.StatusOK, gin.H{
		"flowID":     flowID,
		"csrf_token": csrf_token,
	})
}

func HandlePostMFAFlow(c *gin.Context) {
	var req_body login.SubmitLoginWithMFABody
	err := c.BindJSON(&req_body)

	if err != nil {
		log.ErrorLogger("Unable to process json body", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to process request body",
		})
		return
	}

	flow_cookie, err := c.Cookie("mfa")

	if err != nil {
		log.ErrorLogger("Cookie not found", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "csrf cookie not found",
		})
		return
	}
	session_cookie, _ := c.Cookie("sdslabs_session")
	csrfToken := req_body.CsrfToken
	cookie := strings.Split(flow_cookie, ";")[0] + "; " + strings.Split(session_cookie, ";")[0] + "; x-csrf-token=" + csrfToken

	identity, session, err := login.SubmitLoginWithMFAWrapper(cookie, req_body.FlowID, req_body.CsrfToken, req_body.TOTP)

	if err != nil {
		log.ErrorLogger("Kratos post MFA flow failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.SetCookie("sdslabs_session", session, 3600, "/", config.NymeriaConfig.URL.Domain, true, true)
	c.JSON(http.StatusOK, gin.H{
		"status": "MFA Successful",
		"user":   identity,
	})

}
