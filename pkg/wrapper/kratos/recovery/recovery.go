package recovery

import (
	"context"
	"fmt"
	"os"

	client "github.com/ory/client-go"

	"github.com/sdslabs/nymeria/config"
)

func InitializeRecoveryFlowWrapper() (string, string, string, error) {

	returnTo := "" // string | The URL to return the browser to after the flow was completed. (optional)

	apiClient := client.NewAPIClient(config.KratosClientConfig)

	resp, httpRes, err := apiClient.FrontendAPI.CreateBrowserRecoveryFlow(context.Background()).ReturnTo(returnTo).Execute()
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

	setCookie := httpRes.Header.Get("Set-Cookie")
	return setCookie, resp.Id, csrf_token, nil
}

func SubmitRecoveryFlowWrapper(cookie string, flowID string, csrfToken string, email string) (string, error) {

	submitFlowBody := client.UpdateRecoveryFlowBody{
		UpdateRecoveryFlowWithCodeMethod: client.NewUpdateRecoveryFlowWithCodeMethod("code"),
	}
	submitFlowBody.UpdateRecoveryFlowWithCodeMethod.SetCsrfToken(csrfToken)
	submitFlowBody.UpdateRecoveryFlowWithCodeMethod.SetEmail(email)

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	resp, r, err := apiClient.FrontendAPI.UpdateRecoveryFlow(context.Background()).Flow(flowID).UpdateRecoveryFlowBody(submitFlowBody).Cookie(cookie).Execute()

	var csrf_token string

	for _, node := range resp.Ui.Nodes {
		if node.Attributes.UiNodeInputAttributes.Name == "csrf_token" {
			csrf_token_interface := node.Attributes.UiNodeInputAttributes.Value
			csrf_token, _ = csrf_token_interface.(string)
			break
		}
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.SubmitSelfServiceRecoveryFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return "", err
	}

	return csrf_token, nil
}

func SubmitRecoveryCodeFlowWrapper(cookie string, flowID string, csrfToken string, recoveryCode string) (string, error) {
	submitFlowBody := client.UpdateRecoveryFlowBody{
		UpdateRecoveryFlowWithCodeMethod: client.NewUpdateRecoveryFlowWithCodeMethod("code"),
	}
	submitFlowBody.UpdateRecoveryFlowWithCodeMethod.SetCsrfToken(csrfToken)
	submitFlowBody.UpdateRecoveryFlowWithCodeMethod.SetCode(recoveryCode)

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	_, r, err := apiClient.FrontendAPI.UpdateRecoveryFlow(context.Background()).Flow(flowID).UpdateRecoveryFlowBody(submitFlowBody).Cookie(cookie).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.SubmitSelfServiceRecoveryFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return "", err
	}

	responseCookies := r.Header["Set-Cookie"]

	return responseCookies[1], nil
}
