package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/wrapper/kratos/settings"
)

func HandleGetSettingsFlow(c *gin.Context) {
	log.Logger.Debug("Get Settings")

	auth_cookie, _ := c.Cookie("sdslabs_session")

	cookie, flowID, csrf_token, err := settings.InitializeSettingsFlowWrapper(auth_cookie)

	if err != nil {
		log.ErrorLogger("Intialize Settings Failed", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error": strings.Split(err.Error(), " ")[1],
			"message": "Intialize Settings Failed",
		})
		return
	}

	c.SetCookie("settings_flow", cookie, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"flowID":     flowID,
		"csrf_token": csrf_token,
	})
}

func HandlePostSettingsFlow(c *gin.Context) {
	var t settings.SubmitSettingsWithPasswordBody
	err := c.Bind(&t)

	if err != nil {
		log.ErrorLogger("Unable to process json body", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error": strings.Split(err.Error(), " ")[1],
			"message": "Unable to process json body",
		})
		return
	}

	cookie, err := c.Cookie("settings_flow")

	if err != nil {
		log.ErrorLogger("Cookie not found", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error": strings.Split(err.Error(), " ")[1],
			"message": "Cookie not found",
		})
		return
	}

	_, err = settings.SubmitSettingsFlowWrapper(cookie, t.FlowID, t.CsrfToken, t.Password)

	if err != nil {
		log.ErrorLogger("Post Settings flow failed", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error": strings.Split(err.Error(), " ")[1],
			"message": "Post Settings flow failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password Reset Successful",
	})
}