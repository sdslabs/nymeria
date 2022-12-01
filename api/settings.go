package api

import (
	"net/http"

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
		if node.Attributes.UiNodeInputAttributes.Name == "totp_qr" {
			qr_interface := node.Attributes.UiNodeInputAttributes.Value
			qr, _ = qr_interface.(string)
			break
		}
	}

	var secret string

	for _, node := range flow.Ui.Nodes {
		if node.Attributes.UiNodeInputAttributes.Name == "totp_secret_key" {
			secret_interface := node.Attributes.UiNodeInputAttributes.Value
			secret, _ = secret_interface.(string)
			break
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"flowID":     flowID,
		"csrf_token": csrf_token,
		"qr":         qr,
		"secret":     secret,
	})
}
