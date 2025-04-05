package api

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	client "github.com/ory/client-go"

	"github.com/sdslabs/nymeria/config"
	"github.com/sdslabs/nymeria/helper"
	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/db"
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

// HandleAppAuthorization handles the application authorization request
func HandleAppAuthorization(c *gin.Context) {
	var body SecureAccessProfileRequest
	err := c.BindJSON(&body)
	if err != nil {
		log.ErrorLogger("Unable to process json body", err)
		errCode := helper.ExtractErrorCode(err)
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Unable to process json body",
		})
		return
	}

	// Get the application using only the client key
	app, err := db.GetApplicationByKey(body.ClientKey)
	if err != nil {
		log.ErrorLogger("Unable to get application", err)
		errCode := helper.ExtractErrorCode(err)
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Internal Server Error",
		})
		c.Abort()
		return
	}

	timestampInt, err := strconv.ParseInt(body.Timestamp, 10, 64)
	if err != nil {
		log.ErrorLogger("Invalid timestamp", err)
		c.JSON(400, gin.H{
			"error":   "bad_request",
			"message": "Invalid timestamp",
		})
	}

	// Validate the signature - this proves the client has the secret without sending it
	isValid := helper.ValidateSignature(
		body.ClientKey,
		app.ClientSecret,
		body.RedirectURL,
		body.Signature,
		timestampInt,
	)

	if !isValid {
		log.ErrorLogger("Invalid signature or expired request", nil)
		c.JSON(401, gin.H{
			"error":   "unauthorized",
			"message": "Invalid signature or expired request",
		})
		return
	}

	// Check if redirect URL matches
	if app.RedirectURL != body.RedirectURL {
		log.ErrorLogger("Redirect URL does not match", nil)
		c.JSON(400, gin.H{
			"error":   "bad_request",
			"message": "Redirect URL does not match",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Authorized",
	})
}
