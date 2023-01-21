package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/wrapper/kratos/recovery"
)

func HandleGetRecoveryFlow(c *gin.Context) {
	log.Logger.Debug("Get Recovery")

	cookie, flowID, csrf_token, err := recovery.InitializeRecoveryFlowWrapper()

	if err != nil {
		log.ErrorLogger("Intialize Recovery Failed", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error": strings.Split(err.Error(), " ")[1],
			"message": "Intialize Recovery Failed",
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
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error": strings.Split(err.Error(), " ")[1],
			"message": "Unable to process json body",
		})
		return
	}

	cookie, err := c.Cookie("recovery_flow")

	if err != nil {
		log.ErrorLogger("Cookie not found", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error": strings.Split(err.Error(), " ")[1],
			"message": "Cookie not found",
		})
		return
	}

	_, err = recovery.SubmitRecoveryFlowWrapper(cookie, t.FlowID, t.RecoveryToken, t.CsrfToken, t.Email, t.Method)

	if err != nil {
		log.ErrorLogger("POST Recovery flow failed", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error": strings.Split(err.Error(), " ")[1],
			"message": "POST Recovery flow failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Mail sent with recovery code",
	})
}
