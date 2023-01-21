package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/sdslabs/nymeria/config"
	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/wrapper/kratos/settings"
)

func HandleGetSettingsFlow(c *gin.Context) {
	log.Logger.Debug("Get Settings")
	get_cookie, err := c.Cookie("sdslabs_session")

	if err != nil {
		log.ErrorLogger("Intialize Settings flow Failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Session expired",
		})
		return
	}

	flow, cookie, err := settings.InitializeSettingsFlowWrapper(get_cookie)
	identity := flow.Identity

	flowID := flow.Id

	if err != nil {
		log.ErrorLogger("Intialize Settings flow Failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.SetCookie("settings_flow", cookie, 3600, "/", config.NymeriaConfig.URL.Domain, true, true)

	var csrf_token string

	for _, node := range flow.Ui.Nodes {
		if node.Attributes.UiNodeInputAttributes.Name == "csrf_token" {
			csrf_token_interface := node.Attributes.UiNodeInputAttributes.Value
			csrf_token, _ = csrf_token_interface.(string)
			break
		}
	}

	var qr string

	for _, node := range flow.Ui.Nodes {
		if node.Attributes.UiNodeImageAttributes.Id == "totp_qr" {
			qr = node.Attributes.UiNodeImageAttributes.Src
			break
		}
	}

	var secret string

	for _, node := range flow.Ui.Nodes {
		if node.Attributes.UiNodeTextAttributes.Id == "totp_secret_key" {
			secret = node.Attributes.UiNodeTextAttributes.Text.GetText()
			break
		}
	}
	creds := identity.GetCredentials()

	keys := make([]string, len(creds))
	i := 0
	for k := range creds {
		keys[i] = k
		i++
	}

	for method := range keys {
		if keys[method] == "totp" {
			c.JSON(http.StatusOK, gin.H{
				"status": "unlink button should be visible",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"flowID":     flowID,
		"csrf_token": csrf_token,
		"qr":         qr,
		"secret":     secret,
	})
}

func HandleEnableTOTP(c *gin.Context) {
	var t settings.SubmitSettingsAPIBody
	err := c.BindJSON(&t)

	if err != nil {
		log.ErrorLogger("Unable to process json body", err)

		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Unable to process json body",
		})
		return
	}

	cookie, err := c.Cookie("settings_flow")

	if err != nil {
		log.ErrorLogger("Cookie not found", err)

		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Cookie not found",
		})
		return
	}

	_, session, err := settings.SubmitSettingsFlowWrapper(cookie, t.FlowID, t.CsrfToken, t.TOTPcode, false)

	if err != nil {
		log.ErrorLogger("Kratos post settings flow failed", err)

		errCode, _ := strconv.Atoi((strings.Split(err.Error(), " "))[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Kratos post settings flow failed",
		})
		return
	}

	c.SetCookie("sdslabs_session", session, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"status": "TOTP Enabled",
	})

}

func HandleDisableTOTP(c *gin.Context) {
	var t settings.SubmitSettingsAPIBody
	err := c.BindJSON(&t)

	if err != nil {
		log.ErrorLogger("Unable to process json body", err)

		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Unable to process json body",
		})
		return
	}

	cookie, err := c.Cookie("settings_flow")

	if err != nil {
		log.ErrorLogger("Cookie not found", err)

		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Cookie not found",
		})
		return
	}

	_, session, err := settings.SubmitSettingsFlowWrapper(cookie, t.FlowID, t.CsrfToken, t.TOTPcode, true) // _ is USERID

	if err != nil {
		log.ErrorLogger("Kratos post settings flow failed", err)

		errCode, _ := strconv.Atoi((strings.Split(err.Error(), " "))[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Kratos post settings flow failed",
		})
		return
	}

	c.SetCookie("sdslabs_session", session, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"status": "TOTP Disabled",
	})

}
