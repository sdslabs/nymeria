package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/middleware"
)

func HandlePostProfile(c *gin.Context) {
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

	c.JSON(http.StatusOK, profile)
}
