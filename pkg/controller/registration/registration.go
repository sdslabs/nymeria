package registration

import (
	"context"
	"fmt"
	"os"

	client "github.com/ory/kratos-client-go"
	"github.com/sdslabs/nymeria/config"
)

func InitializeRegistrationFlowWrapper() (string, string, string, error) {
	returnTo := "http://127.0.0.1:4455/ping"

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	resp, r, err := apiClient.V0alpha2Api.InitializeSelfServiceRegistrationFlowForBrowsers(context.Background()).ReturnTo(returnTo).Execute()
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

func SubmitRegistrationFlowWrapper(cookie string, flowID string, csrfToken string, password string, data Traits) (string, error) {
	trait := map[string]interface{}{
		"email": data.Email,
		"name":  data.Name,
	}

	submitDataBody := client.SubmitSelfServiceRegistrationFlowBody{
		SubmitSelfServiceRegistrationFlowWithPasswordMethodBody: client.NewSubmitSelfServiceRegistrationFlowWithPasswordMethodBody("password", password, trait),
	}

	submitDataBody.SubmitSelfServiceRegistrationFlowWithPasswordMethodBody.SetCsrfToken(csrfToken)

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	_, r, err := apiClient.V0alpha2Api.SubmitSelfServiceRegistrationFlow(context.Background()).Flow(flowID).SubmitSelfServiceRegistrationFlowBody(submitDataBody).Cookie(cookie).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.SubmitSelfServiceRegistrationFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseCookies := r.Header["Set-Cookie"]
	return responseCookies[1], nil
}