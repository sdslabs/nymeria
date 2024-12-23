package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sdslabs/nymeria/config"
	"github.com/sdslabs/nymeria/helper"
	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/wrapper/kratos/recovery"
)

func HandleGetRecoveryFlow(c *gin.Context) {
	log.Logger.Debug("Get Recovery")

	cookie, flowID, csrf_token, err := recovery.InitializeRecoveryFlowWrapper()

	if err != nil {
		log.ErrorLogger("Initialize Recovery Failed", err)
		errCode := helper.ExtractErrorCode(err)
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Initialize Recovery Failed",
		})
		return
	}

	c.SetCookie("recovery_flow", cookie, 3600, "/", config.NymeriaConfig.URL.Domain, true, true)

	c.JSON(http.StatusOK, gin.H{
		"flowID":     flowID,
		"csrf_token": csrf_token,
	})
}

func HandlePostRecoveryFlow(c *gin.Context) {
	var t recovery.SubmitRecoveryAPIBody
	err := c.BindJSON(&t)

	if err != nil {
		log.ErrorLogger("Unable to process json body", err)
		errCode := helper.ExtractErrorCode(err)
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Unable to process json body",
		})
		return
	}

	cookie, err := c.Cookie("recovery_flow")

	if err != nil {
		log.ErrorLogger("Cookie not found", err)
		errCode := helper.ExtractErrorCode(err)
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Cookie not found",
		})
		return
	}
	var csrf_token string
	csrf_token, err = recovery.SubmitRecoveryFlowWrapper(cookie, t.FlowID, t.CsrfToken, t.Email)

	if err != nil {
		log.ErrorLogger("POST Recovery flow failed", err)
		errCode := helper.ExtractErrorCode(err)
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "POST Recovery flow failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Mail sent with recovery code",
		"csrf_token": csrf_token,
	})
}

func HandlePostRecoveryCodeFlow(c *gin.Context) {
	var t recovery.SubmitRecoveryAPIBody
	err := c.BindJSON(&t)

	if err != nil {
		log.ErrorLogger("Unable to process json body", err)
		errCode := helper.ExtractErrorCode(err)
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Unable to process json body",
		})
		return
	}

	cookie, err := c.Cookie("recovery_flow")

	if err != nil {
		log.ErrorLogger("Recovery Flow Cookie not found", err)
		errCode := helper.ExtractErrorCode(err)
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Recovery Flow Cookie not found",
		})
		return
	}

	session, err := recovery.SubmitRecoveryCodeFlowWrapper(cookie, t.FlowID, t.CsrfToken, t.RecoveryCode)

	if err != nil {
		log.ErrorLogger("POST Recovery flow failed", err)
		errCode := helper.ExtractErrorCode(err)
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "POST Recovery Code flow failed",
		})
		return
	}

	c.SetCookie("sdslabs_session", session, 3600, "/", config.NymeriaConfig.URL.Domain, true, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Session cookie set for password change",
	})
}
