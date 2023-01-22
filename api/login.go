package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/sdslabs/nymeria/config"
	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/wrapper/kratos/login"
)

func HandleGetLoginFlow(c *gin.Context) {
	log.Logger.Debug("Get Login")
	cookie, flowID, csrf_token, err := login.InitializeLoginFlowWrapper("aal1")

	if err != nil {
		log.ErrorLogger("Initialize Login Failed", err)

		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Initialize Login Failed",
		})
		return
	}

	c.SetCookie("login_flow", cookie, 3600, "/", config.NymeriaConfig.URL.Domain, true, true)

	c.JSON(http.StatusOK, gin.H{
		"flowID":     flowID,
		"csrf_token": csrf_token,
	})
}

func HandlePostLoginFlow(c *gin.Context) {
	var t login.SubmitLoginAPIBody
	err := c.BindJSON(&t)

	if err != nil {
		log.ErrorLogger("Unable to process json body", err)

		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Unable to process json body",
		})
		return
	}

	cookie, err := c.Cookie("login_flow")

	if err != nil {
		log.ErrorLogger("Cookie not found", err)

		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Cookie not found",
		})
		return
	}

	identity, session, err := login.SubmitLoginFlowWrapper(cookie, t.FlowID, t.CsrfToken, t.Password, t.Identifier) // _ is USERID

	if err != nil {
		log.ErrorLogger("Post login flow failed", err)

		errCode, _ := strconv.Atoi((strings.Split(err.Error(), " "))[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Kratos post login flow failed",
		})
		return
	}

	c.SetCookie("sdslabs_session", session, 3600, "/", config.NymeriaConfig.URL.Domain, true, true)
	c.JSON(http.StatusOK, gin.H{
		"status": "user logged in",
		"person": identity,
	})

}
