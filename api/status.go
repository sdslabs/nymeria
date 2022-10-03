package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	client "github.com/ory/kratos-client-go"
	"github.com/sdslabs/nymeria/config"
)

func HandleStatus(c *gin.Context) {
	cookie, err := c.Cookie("sdslabs_session")

	if err != nil {
		fmt.Println(err)

	}

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	resp, _, err := apiClient.V0alpha2Api.ToSession(context.Background()).Cookie(cookie).Execute()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
	fmt.Println(cookie)
	c.JSON(http.StatusOK, gin.H{
		"flowID": resp.Active,
	})
}
