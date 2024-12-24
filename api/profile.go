package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sdslabs/nymeria/helper"
	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/middleware"
)

func HandlePostProfile(c *gin.Context) {
	session, err := middleware.GetSession(c)
	if err != nil {
		log.ErrorLogger("Unable to get session", err)
		errCode := helper.ExtractErrorCode(err)
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Unable to get session",
		})
		return
	}
	identity := session.GetIdentity()
	traits := identity.GetTraits()
	profile := traits.(map[string]interface{})

	c.JSON(http.StatusOK, profile)
}

func HandleGetVerifiedStatus(c *gin.Context) {
	session, err := middleware.GetSession(c)
	if err != nil {
		log.ErrorLogger("Unable to get session", err)
		errCode := helper.ExtractErrorCode(err)
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Unable to get session",
		})
		return
	}
	identity := session.GetIdentity()
	verifiableAddresses := identity.GetVerifiableAddresses()
	emails := verifiableAddresses

	c.JSON(http.StatusOK, emails)
}
