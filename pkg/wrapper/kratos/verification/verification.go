package verification

import (
	"context"
	"fmt"
	"os"

	client "github.com/ory/client-go"

	"github.com/sdslabs/nymeria/config"
)

func InitializeVerificationFlowWrapper() (string, string, string, error) {
	apiClient := client.NewAPIClient(config.KratosClientConfig)

	resp, r, err := apiClient.V0alpha2Api.InitializeSelfServiceVerificationFlowForBrowsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.InitializeSelfServiceVerificationFlowForBrowsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
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

	setCookie := r.Header.Get("Set-Cookie")
	return setCookie, resp.Id, csrf_token, nil
}

func SubmitVerificationFlowWrapper(cookie string, flowID string, csrfToken string, email string) (string, error) {

	submitFlowBody := client.SubmitSelfServiceVerificationFlowBody{
		SubmitSelfServiceVerificationFlowWithLinkMethodBody: client.NewSubmitSelfServiceVerificationFlowWithLinkMethodBody(email, "link"),
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
