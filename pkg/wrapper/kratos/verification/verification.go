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

	resp, r, err := apiClient.FrontendAPI.CreateBrowserVerificationFlow(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.InitializeSelfServiceVerificationFlowForBrowsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	var csrf_token string

	for _, node := range resp.Ui.Nodes {
		if node.Attributes.UiNodeInputAttributes.Name == "csrf_token" {
			csrf_token_interface := node.Attributes.UiNodeInputAttributes.Value
			csrf_token, _ = csrf_token_interface.(string)
			break
		}
	}

	setCookie := r.Header.Get("Set-Cookie")
	return setCookie, resp.Id, csrf_token, nil
}

func InitializeVerificationAfterRegistrationFlowWrapper(cookie string, flowID string) (string, error) {
	apiClient := client.NewAPIClient(config.KratosClientConfig)

	resp, r, err := apiClient.FrontendAPI.GetVerificationFlow(context.Background()).Cookie(cookie).Id(flowID).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.InitializeSelfServiceVerificationFlowForBrowsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	var csrf_token string

	for _, node := range resp.Ui.Nodes {
		if node.Attributes.UiNodeInputAttributes.Name == "csrf_token" {
			csrf_token_interface := node.Attributes.UiNodeInputAttributes.Value
			csrf_token, _ = csrf_token_interface.(string)
			break
		}
	}

	return csrf_token, nil
}

func SubmitVerificationFlowWrapper(cookie string, flowID string, csrfToken string, email string) (string, error) {

	submitFlowBody := client.UpdateVerificationFlowBody{
		UpdateVerificationFlowWithCodeMethod: client.NewUpdateVerificationFlowWithCodeMethod("code"),
	}

	submitFlowBody.UpdateVerificationFlowWithCodeMethod.SetCsrfToken(csrfToken)
	submitFlowBody.UpdateVerificationFlowWithCodeMethod.SetEmail(email)

	apiClient := client.NewAPIClient(config.KratosClientConfig)

	resp, r, err := apiClient.FrontendAPI.UpdateVerificationFlow(context.Background()).Flow(flowID).UpdateVerificationFlowBody(submitFlowBody).Cookie(cookie).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.SubmitSelfServiceVerificationFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return "", err
	}

	var csrf_token string

	for _, node := range resp.Ui.Nodes {
		if node.Attributes.UiNodeInputAttributes.Name == "csrf_token" {
			csrf_token_interface := node.Attributes.UiNodeInputAttributes.Value
			csrf_token, _ = csrf_token_interface.(string)
			break
		}
	}

	return csrf_token, nil
}

func SubmitVerificationCodeFlowWrapper(cookie string, flowID string, csrfToken string, verificationCode string) (string, error) {

	submitFlowBody := client.UpdateVerificationFlowBody{
		UpdateVerificationFlowWithCodeMethod: client.NewUpdateVerificationFlowWithCodeMethod("code"),
	}

	submitFlowBody.UpdateVerificationFlowWithCodeMethod.SetCsrfToken(csrfToken)
	submitFlowBody.UpdateVerificationFlowWithCodeMethod.SetCode(verificationCode)

	apiClient := client.NewAPIClient(config.KratosClientConfig)

	_, r, err := apiClient.FrontendAPI.UpdateVerificationFlow(context.Background()).Flow(flowID).Token(verificationCode).UpdateVerificationFlowBody(submitFlowBody).Cookie(cookie).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.SubmitSelfServiceVerificationFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return "", err
	}

	return "", nil
}
