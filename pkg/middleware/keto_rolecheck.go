package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/wrapper/keto"
)

func CheckIfAllowed(c *gin.Context) {
	session, err := GetSession(c)
	if err != nil {
		log.ErrorLogger("Couldn't retrieve session: ", err)
		c.Abort()
		return
	}
	identity := session.GetIdentity()
	traits := identity.GetTraits()
	role := traits.(map[string]interface{})["role"]

	requestedRoute := c.Request.URL.String()

	data := map[string]interface{}{
		"namespace":  "accounts",
		"object":     requestedRoute,
		"relation":   "view",
		"subject_id": role,
	}

	response, err := keto.MakeRequest(keto.CheckPermissionEndpoint, data)
	if err != nil {
		log.ErrorLogger("Error in making request to keto", err)
		c.Abort()
		return
	}

	if response["allowed"] == true {
		c.Next()
		return
	} else {
		c.JSON(403, gin.H{
			"error":   "Forbidden",
			"message": "You don't have permission to access this resource.",
		})
		c.Abort()
		return
	}
}
