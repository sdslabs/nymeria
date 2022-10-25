package recovery

import (
	"context"

	client "github.com/ory/kratos-client-go"
	"github.com/sdslabs/nymeria/config"
)

func InitializeRecoveryFlowWrapper() (string, string, string, error){

	returnTo := "http://127.0.0.1:4455/ping" // string | The URL to return the browser to after the flow was completed. (optional)

	apiClient := client.NewAPIClient(config.KratosClientConfig)

	resp, httpRes, err := apiClient.V0alpha2Api.InitializeSelfServiceRecoveryFlowForBrowsers(context.Background()).ReturnTo(returnTo).Execute()
	if err != nil {
		return "", "", "", err
	}

	cookie := httpRes.Header.Get("Host")

	var csrf_token string

	for _, node := range resp.Ui.Nodes {
		if node.Attributes.UiNodeInputAttributes.Name == "csrf_token" {
			csrf_token_interface := node.Attributes.UiNodeInputAttributes.Value
			csrf_token, _ = csrf_token_interface.(string)
			break
		}
	}

	return cookie, resp.Id, csrf_token, nil
}

func SubmitRecoveryFlowWrapper(cookie string, flowID string, csrfToken string, password string, data Traits) (string, error) {
	//apiClient := client.NewAPIClient(config.KratosClientConfig)

	return "",nil
}
