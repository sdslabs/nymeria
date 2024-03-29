package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/sdslabs/nymeria/config"
	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/middleware"
	"github.com/sdslabs/nymeria/pkg/wrapper/kratos/settings"
)

func HandleGetSettingsFlow(c *gin.Context) {
	log.Logger.Debug("Get Settings")
	session_cookie, err1 := c.Cookie("sdslabs_session")
	recovery_cookie, err2 := c.Cookie("ory_kratos_session")

	if err1 != nil && err2 != nil {
		var err error
		if err1 != nil {
			err = err1
		} else {
			err = err2
		}
		log.ErrorLogger("Initialize Settings Failed", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Initialize Settings Failed",
		})
		return
	}

	flow, flow_cookie, err := settings.InitializeSettingsFlowWrapper(session_cookie, recovery_cookie)

	if err != nil {
		log.ErrorLogger("Initialize Settings flow Failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.SetCookie("settings_flow", flow_cookie, 3600, "/", config.NymeriaConfig.URL.Domain, true, true)

	if recovery_cookie != "" {
		recovery_cookie = "ory_kratos_session=" + recovery_cookie
		c.SetCookie("sdslabs_session", recovery_cookie, 900, "/", config.NymeriaConfig.URL.Domain, false, true)
	}

	flowID := flow.GetId()

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

	session, err := middleware.GetSession(c)
	if err != nil {
		log.ErrorLogger("Unable to get session", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Unable to get session",
		})
		return
	}
	identity := session.GetIdentity()
	traits := identity.GetTraits().(map[string]interface{})

	if identity.VerifiableAddresses[0].Verified && traits["verified"] == false {
		traits["verified"] = true

		_, err = settings.SubmitSettingsFlowProfileMethod(flow_cookie, session_cookie, flowID, csrf_token, traits)
		if err != nil {
			log.ErrorLogger("Kratos post settings update profile flow failed", err)

			errCode, _ := strconv.Atoi((strings.Split(err.Error(), " "))[0])
			c.JSON(errCode, gin.H{
				"error":   err.Error(),
				"message": "Kratos post settings update profile flow failed",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"flowID":      flowID,
		"csrf_token":  csrf_token,
		"qr":          qr_src,
		"totp_secret": totp_secret,
	})
}

func HandleUpdateProfile(c *gin.Context) {
	var req_body settings.UpdateProfileAPIBody
	err := c.BindJSON(&req_body)

	traitsinterface := map[string]interface{}{
		"email":         req_body.Traits.Email,
		"name":          req_body.Traits.Name,
		"password":      req_body.Traits.Password,
		"img_url":       req_body.Traits.ImgURL,
		"phone_number":  req_body.Traits.PhoneNumber,
		"invite_status": req_body.Traits.InviteStatus,
		"verified":      req_body.Traits.Verified,
		"role":          req_body.Traits.Role,
		"created_at":    req_body.Traits.Created_At,
		"totp_enabled":  req_body.Traits.TOTP_Enabled,
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

	//Checking if email is changed then verified will be false
	session, err := middleware.GetSession(c)
	if err != nil {
		log.ErrorLogger("Unable to get session", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Unable to get session",
		})
		return
	}
	identity := session.GetIdentity()
	traits := identity.GetTraits()
	profile := traits.(map[string]interface{})

	if profile["email"] != traitsinterface["email"] {
		traitsinterface["verified"] = false
	}

	msg, err := settings.SubmitSettingsFlowProfileMethod(flow_cookie, session_cookie, req_body.FlowID, req_body.CsrfToken, traitsinterface)

	if err != nil {
		log.ErrorLogger("Kratos post settings update profile flow failed", err)

		errCode, _ := strconv.Atoi((strings.Split(err.Error(), " "))[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Kratos post settings update profile flow failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": msg,
	})
}

func HandleChangePassword(c *gin.Context) {
	var req_body settings.ChangePasswordAPIBody
	err := c.BindJSON(&req_body)

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

	msg, err := settings.SubmitSettingsFlowPasswordMethod(flow_cookie, session_cookie, req_body.FlowID, req_body.CsrfToken, req_body.Password)

	if err != nil {
		log.ErrorLogger("Kratos post settings change password flow failed", err)

		errCode, _ := strconv.Atoi((strings.Split(err.Error(), " "))[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Kratos post settings change password flow failed",
		})
		return
	}

	recovery_cookie, _ := c.Cookie("ory_kratos_session")

	session, err := middleware.GetSession(c)
	if err != nil {
		log.ErrorLogger("Unable to get session", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Unable to get session",
		})
		return
	}
	identity := session.GetIdentity()
	traits := identity.GetTraits()
	profile := traits.(map[string]interface{})

	profile["password"] = req_body.Password

	if recovery_cookie != "" {
		if profile["verified"] == false {
			profile["verified"] = true
		}
		if profile["invite_status"] == "pending" {
			profile["invite_status"] = "accepted"
		}
	}

	_, err = settings.SubmitSettingsFlowProfileMethod(flow_cookie, session_cookie, req_body.FlowID, req_body.CsrfToken, profile)
	if err != nil {
		log.ErrorLogger("Kratos post settings update profile flow failed", err)

		errCode, _ := strconv.Atoi((strings.Split(err.Error(), " "))[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Kratos post settings update profile flow failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": msg,
	})
}

func HandleToggleTOTP(c *gin.Context) {
	var req_body settings.ToggleTOTPAPIBody
	err := c.BindJSON(&req_body)

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

	msg, err := settings.SubmitSettingsFlowTOTPMethod(flow_cookie, session_cookie, req_body.FlowID, req_body.CsrfToken, req_body.TOTPCode, req_body.TOTPUnlink)

	if err != nil {
		log.ErrorLogger("Kratos post settings toggle totp flow failed", err)

		errCode, _ := strconv.Atoi((strings.Split(err.Error(), " "))[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Kratos post settings toggle totp flow failed",
		})
		return
	}

	// Updates traits in identity
	session, err := middleware.GetSession(c)
	if err != nil {
		log.ErrorLogger("Unable to get session", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Unable to get session",
		})
		return
	}
	identity := session.GetIdentity()
	profile := identity.GetTraits().(map[string]interface{})

	if !req_body.TOTPUnlink {
		profile["totp_enabled"] = true
	} else {
		profile["totp_enabled"] = false
	}

	_, err = settings.SubmitSettingsFlowProfileMethod(flow_cookie, session_cookie, req_body.FlowID, req_body.CsrfToken, profile)
	if err != nil {
		log.ErrorLogger("Kratos post settings update profile flow failed", err)

		errCode, _ := strconv.Atoi((strings.Split(err.Error(), " "))[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Kratos post settings update profile flow failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": msg,
	})
}
