package middleware

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/sdslabs/nymeria/helper"
	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/db"
)

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
		c.Abort()
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
		c.Abort()
		return
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
		c.Abort()
		return
	}

	// Check if redirect URL matches
	if app.RedirectURL != body.RedirectURL {
		log.ErrorLogger("Redirect URL does not match", nil)
		c.JSON(400, gin.H{
			"error":   "bad_request",
			"message": "Redirect URL does not match",
		})
		c.Abort()
		return
	}

	c.Next()
}
