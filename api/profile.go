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

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile fetched successfully",
		"profile": gin.H{
			"identityId": identity.GetId(),
			"status":     identity.GetState(),
			"traits":     traits,
		},
	})
}
