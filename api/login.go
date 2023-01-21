package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/wrapper/kratos/login"
)

func HandleGetLoginFlow(c *gin.Context) {
	log.Logger.Debug("Get Login")
	cookie, flowID, csrf_token, err := login.InitializeLoginFlowWrapper()

	if err != nil {
		log.ErrorLogger("Intialize Login Failed", err)

		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error": strings.Split(err.Error(), " ")[1],
			"message": "Intialize Login Failed",
		})
		return
	}

	c.SetCookie("login_flow", cookie, 3600, "/", "localhost", false, true)

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
			"error": strings.Split(err.Error(), " ")[1],
			"message": "Unable to process json body",
		})
		return
	}

	cookie, err := c.Cookie("login_flow")

	if err != nil {
		log.ErrorLogger("Cookie not found", err)

		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error": strings.Split(err.Error(), " ")[1],
			"message": "Cookie not found",
		})
		return
	}

	session, err := login.SubmitLoginFlowWrapper(cookie, t.FlowID, t.CsrfToken, t.Password, t.Identifier)

	if err != nil {
		log.ErrorLogger("Kratos post login flow failed", err)

		errCode, _ := strconv.Atoi((strings.Split(err.Error(), " "))[0])
		c.JSON(errCode, gin.H{
			"error": strings.Split(err.Error(), " ")[1],
			"message": "Kratos post login flow failed",
		})
		return
	}

	c.SetCookie("sdslabs_session", session, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"status": "user logged in",
	})

}
