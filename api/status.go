package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	client "github.com/ory/client-go"

	"github.com/sdslabs/nymeria/config"
	"github.com/sdslabs/nymeria/helper"
	"github.com/sdslabs/nymeria/log"
)

// HandleVerifySession handles the user session verification request
func HandleVerifySession(c *gin.Context) {
	cookie, err := c.Cookie("sdslabs_session")
	if err != nil {
		log.ErrorLogger("Session cookie not found", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"message": "Session cookie not found",
		})
		return
	}

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	resp, r, err := apiClient.FrontendAPI.ToSession(context.Background()).Cookie(cookie).Execute()
	if err != nil {
		msg := helper.ExtractErrorMessage(r)
		errCode := helper.ExtractErrorCode(err)

		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": msg,
		})
		return
	}

	if !*resp.Active {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Session expired",
			"message": "Session expired",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Session verified",
	})
}
