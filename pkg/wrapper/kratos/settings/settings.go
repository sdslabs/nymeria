package settings

import (
	"context"
	"fmt"
	"os"

	client "github.com/ory/client-go"
	"github.com/sdslabs/nymeria/config"
)

func InitializeSettingsFlowWrapper(auth_cookie string) (string, string, string, error) {
	returnTo := "http://127.0.0.1:4455/ping"

	apiClient := client.NewAPIClient(config.KratosClientConfig)

	resp, r, err := apiClient.V0alpha2Api.InitializeSelfServiceSettingsFlowForBrowsers(context.Background()).ReturnTo(returnTo).Cookie(auth_cookie).Execute()

	if err != nil {
		return "", "", "", err
	}

	var csrf_token string

	for _, node := range resp.Ui.Nodes {

		if node.Attributes.UiNodeInputAttributes.Name == "csrf_token" {
			csrf_token_interface := node.Attributes.UiNodeInputAttributes.Value
			csrf_token, _ = csrf_token_interface.(string)
			break
		}
	}

	var setCookie string = r.Header.Get("Set-Cookie")
	return setCookie, resp.Id, csrf_token, nil
}

func SubmitSettingsFlowWrapper(cookie string, session string, flowID string, csrfToken string, pass string) (string, error) {
	submitDataBody := client.SubmitSelfServiceSettingsFlowBody{
		SubmitSelfServiceSettingsFlowWithPasswordMethodBody: client.NewSubmitSelfServiceSettingsFlowWithPasswordMethodBody("password", pass)}

	submitDataBody.SubmitSelfServiceSettingsFlowWithPasswordMethodBody.SetCsrfToken(csrfToken)

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	_, r, err := apiClient.V0alpha2Api.SubmitSelfServiceSettingsFlow(context.Background()).Flow(flowID).SubmitSelfServiceSettingsFlowBody(submitDataBody).XSessionToken(session).Cookie(cookie).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.SubmitSelfServiceSettingsFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return "", err
	}

	return "", nil
}
