package verification

import (
	"context"
	"fmt"
	"os"

	client "github.com/ory/client-go"
	"github.com/sdslabs/nymeria/config"
)

func InitializeVerificationFlowWrapper(auth_cookie string) (string, string, string, error) {
	returnTo := "http://127.0.0.1:4455/ping" // string | The URL to return the browser to after the flow was completed. (optional)

	apiClient := client.NewAPIClient(config.KratosClientConfig)

	resp, httpRes, err := apiClient.V0alpha2Api.InitializeSelfServiceRecoveryFlowForBrowsers(context.Background()).ReturnTo(returnTo).Execute()
	if err != nil {
		return "", "", "", err
	}

	var csrf_token string

	for _, node := range resp.Ui.Nodes {
		fmt.Println(node.Attributes.UiNodeInputAttributes)
		if node.Attributes.UiNodeInputAttributes.Name == "csrf_token" {
			csrf_token_interface := node.Attributes.UiNodeInputAttributes.Value
			csrf_token, _ = csrf_token_interface.(string)
			break
		}
	}

	setCookie := httpRes.Header.Get("Set-Cookie")
	return setCookie, resp.Id, csrf_token, nil
}

func SubmitVerificationFlowWrapper(cookie string, session string, flowID string, csrfToken string, email string, method string) (string, error) {

	submitFlowBody := client.SubmitSelfServiceVerificationFlowBody{
		SubmitSelfServiceVerificationFlowWithLinkMethodBody: client.NewSubmitSelfServiceVerificationFlowWithLinkMethodBody(email, method),
	}

	submitFlowBody.SubmitSelfServiceVerificationFlowWithLinkMethodBody.SetCsrfToken(csrfToken)

	apiClient := client.NewAPIClient(config.KratosClientConfig)

	_, r, err := apiClient.V0alpha2Api.SubmitSelfServiceVerificationFlow(context.Background()).Flow(flowID).SubmitSelfServiceVerificationFlowBody(submitFlowBody).Cookie(cookie).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.SubmitSelfServiceVerificationFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return "", err
	}

	return "", nil
}
