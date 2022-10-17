package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	client "github.com/ory/kratos-client-go"
	"github.com/sdslabs/nymeria/config"
	"github.com/sdslabs/nymeria/log"
)

func HandleStatus(c *gin.Context) {
	cookie, err := c.Cookie("sdslabs_session")

	if err != nil {
		log.ErrorLogger("Session cookie not found", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "session not found",
		})
		return
	}

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	resp, _, err := apiClient.V0alpha2Api.ToSession(context.Background()).Cookie(cookie).Execute()
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"flowID": resp.Active,
	})
}
