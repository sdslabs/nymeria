package middleware

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	client "github.com/ory/client-go"

	"github.com/sdslabs/nymeria/config"
	"github.com/sdslabs/nymeria/helper"
	"github.com/sdslabs/nymeria/log"
)

type kratosMiddleware struct {
	client *client.APIClient
}

func NewMiddleware() *kratosMiddleware {
	configuration := client.NewConfiguration()
	configuration.Servers = []client.ServerConfiguration{
		{
			URL: "http://127.0.0.1:4433", // Kratos Admin API
		},
	}
	return &kratosMiddleware{
		client: client.NewAPIClient(configuration),
	}
}

func NewAdminMiddleware() *client.APIClient {
	configuration := client.NewConfiguration()
	configuration.Servers = []client.ServerConfiguration{
		{
			URL: "http://127.0.0.1:4434", // Kratos Public API
		},
	}

	apiClient := client.NewAPIClient(configuration)
	return apiClient
}

func GetSession(c *gin.Context) (*client.Session, error) {
	cookie, err := c.Cookie("sdslabs_session")
	if err != nil {
		log.ErrorLogger("Session cookie not found", err)
		return nil, err
	}
	apiClient := client.NewAPIClient(config.KratosClientConfig)
	resp, r, err := apiClient.FrontendAPI.ToSession(context.Background()).Cookie(cookie).Execute()
	if err != nil {
		msg := helper.ExtractErrorMessage(r)
		log.ErrorLogger("Error when calling `FrontendAPI.ToSession`", err)
		return nil, errors.New(msg)
	}
	return resp, nil
}
