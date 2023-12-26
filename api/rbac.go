package api

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	client "github.com/ory/client-go"

	"github.com/sdslabs/nymeria/config"
	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/wrapper/keto"
)

func HandleRbac(c *gin.Context) {
	log.Logger.Debug("RBAC")
	cookie, err := c.Cookie("sdslabs_session")

	if err != nil {
		log.ErrorLogger("Session cookie not found", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Initialize Rbac failed.",
		})
		return
	}

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	session, _, err := apiClient.V0alpha2Api.ToSession(context.Background()).Cookie(cookie).Execute()
	if err != nil {
		log.ErrorLogger("Invalid Cookie", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Initialize Rbac failed.",
		})
		return
	}

	identity := session.GetIdentity()
	traits := identity.GetTraits()
	role := traits.(map[string]interface{})["role"]

	data := map[string]interface{}{
		"namespace": "accounts",
		"relation": "view",
		"subject_id": role,
	}

	res, err := keto.MakeRequest(keto.QueryRelationshipsEndpoint, data)
	if err != nil {
		log.ErrorLogger("Error in making request to keto", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Creating relationship failed.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "RBAC passed",
		"role":    role,
		"res":     res,
	})
}
