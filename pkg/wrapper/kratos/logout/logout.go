package logout

import (
	"context"

	client "github.com/ory/client-go"

	"github.com/sdslabs/nymeria/config"
)

func InitializeLogoutFlowWrapper(cookie string) (*client.LogoutFlow, error) {
	apiClient := client.NewAPIClient(config.KratosClientConfig)

	resp, _, err := apiClient.FrontendAPI.CreateBrowserLogoutFlow(context.Background()).Cookie(cookie).Execute()

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func SubmitLogoutFlowWrapper(cookie string, token string, returnToUrl string) error {
	apiClient := client.NewAPIClient(config.KratosClientConfig)

	_, err := apiClient.FrontendAPI.UpdateLogoutFlow(context.Background()).Cookie(cookie).Token(token).Execute()
	if err != nil {
		return err
	}

	return nil
}
