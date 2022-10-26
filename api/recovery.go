package api

import (
	"net/http"

	"github.com/sdslabs/nymeria/pkg/wrapper/kratos/recovery"
	"github.com/sdslabs/nymeria/log"
	"github.com/gin-gonic/gin"
)

func HandleGetRecoveryFlow(c *gin.Context) {
	log.Logger.Debug("Get Recovery")
	
	cookie, flowID, csrf_token, err := recovery.InitializeRecoveryFlowWrapper()

	if err != nil {
		log.ErrorLogger("Intialize Recovery Failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.SetCookie("recovery_flow", cookie, 3600, "/", "localhost", false, true)

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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to process request body",
		})
		return
	}

	cookie, err := c.Cookie("recovery_flow")

	if err != nil {
		log.ErrorLogger("Cookie not found", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "csrf cookie not found",
		})
		return
	}

	recoveryToken, err := recovery.SubmitRecoveryFlowWrapper(cookie, t.FlowID, t.RecoveryToken, t.CsrfToken, t.Email, t.Method)

	if err != nil {
		log.ErrorLogger("POST Recovery flow failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token" : recoveryToken,
	})
}