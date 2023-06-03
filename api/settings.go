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
	session_cookie, err := c.Cookie("sdslabs_session")

	if err != nil {
		log.ErrorLogger("Initialize Settings Failed", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Initialize Settings Failed",
		})
		return
	}

	flow, flow_cookie, err := settings.InitializeSettingsFlowWrapper(session_cookie)

	c.SetCookie("settings_flow", flow_cookie, 3600, "/", config.NymeriaConfig.URL.Domain, false, true)

	flowID := flow.GetId()

	if err != nil {
		log.ErrorLogger("Intialize Settings flow Failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	var csrf_token string

	for _, node := range flow.Ui.Nodes {
		if node.Attributes.UiNodeInputAttributes.Name == "csrf_token" {
			csrf_token_interface := node.Attributes.UiNodeInputAttributes.GetValue()
			csrf_token, _ = csrf_token_interface.(string)
			break
		}
	}

	var qr_src string

	for _, node := range flow.Ui.Nodes {
		if node.Attributes.UiNodeImageAttributes.GetId() == "totp_qr" {
			qr_src = node.Attributes.UiNodeImageAttributes.GetSrc()
			break
		}
	}

	var totp_secret string

	for _, node := range flow.Ui.Nodes {
		if node.Attributes.UiNodeTextAttributes.GetId() == "totp_secret_key" {
			totp_secret = node.Attributes.UiNodeTextAttributes.Text.GetText()
			break
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"flowID":      flowID,
		"csrf_token":  csrf_token,
		"qr":          qr_src,
		"totp_secret": totp_secret,
	})
}

func HandlePostSettingsFlow(c *gin.Context) {
	var req_body settings.SubmitSettingsAPIBody
	err := c.BindJSON(&req_body)

	traitsinterface := map[string]interface{}{
		"name":         req_body.Traits.Name,
		"phone_number": req_body.Traits.PhoneNumber,
		"img_url":      req_body.Traits.ImgURL,
		"email":        req_body.Traits.Email,
		"totp_enabled": req_body.Traits.TOTP_Enabled,
		"created_at":   req_body.Traits.Created_At,
		"role":         req_body.Traits.Role,
		"active":       req_body.Traits.Active,
	}

	if err != nil {
		log.ErrorLogger("Unable to process json body", err)

		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Unable to process json body",
		})
		return
	}

	flow_cookie, err := c.Cookie("settings_flow")
	if err != nil {
		log.ErrorLogger("Flow Cookie not found", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Cookie not found",
		})
		return
	}

	session_cookie, err := c.Cookie("sdslabs_session")
	if err != nil {
		log.ErrorLogger("Session Cookie not found", err)

		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Cookie not found",
		})
		return
	}

	msg, err := settings.SubmitSettingsFlowWrapper(flow_cookie, session_cookie, req_body.FlowID, req_body.CsrfToken, req_body.Method, req_body.TOTPCode, req_body.TOTPUnlink, req_body.Password, traitsinterface)

	if err != nil {
		log.ErrorLogger("Kratos post settings flow failed", err)

		errCode, _ := strconv.Atoi((strings.Split(err.Error(), " "))[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Kratos post settings flow failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": msg,
	})
}
