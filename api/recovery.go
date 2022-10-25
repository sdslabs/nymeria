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
	
}