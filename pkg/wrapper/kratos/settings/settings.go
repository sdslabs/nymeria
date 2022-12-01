package settings

import (
	"context"

	client "github.com/ory/client-go"

	"github.com/sdslabs/nymeria/config"
)

func InitializeSettingsFlowWrapper(req_cookie string) (client.SelfServiceSettingsFlow, string, error) {

	returnTo := "http://127.0.0.1:4455/ping" // string | The URL to return the browser to after the flow was completed. (optional)

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	resp, httpRes, err := apiClient.V0alpha2Api.InitializeSelfServiceSettingsFlowForBrowsers(context.Background()).ReturnTo(returnTo).Cookie(req_cookie).Execute()
	if err != nil {
		return *resp, "", err
	}

	cookie := httpRes.Header.Get("Host")

	return *resp, cookie, nil
}
