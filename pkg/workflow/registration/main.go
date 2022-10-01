package registration

import (
	"context"
	"fmt"
	"os"

	client "github.com/ory/kratos-client-go"
)

func InitializeRegistrationFlowWrapper() (string, string, string, error) {
	returnTo := "http://127.0.0.1:4455/ping"

	configuration := client.NewConfiguration()
	configuration.Servers = []client.ServerConfiguration{
		{
			URL: "http://127.0.0.1:4433",
		},
	}

	apiClient := client.NewAPIClient(configuration)
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

func SubmitRegistrationFlowWrapper(cookie string, flowID string, csrfToken string, pass string, data Traits) error {
	trait := map[string]interface{}{
		"email": data.Email,
		"name":  data.Name,
	}

	configuration := client.NewConfiguration()
	configuration.Servers = []client.ServerConfiguration{
		{
			URL: "http://127.0.0.1:4433",
		},
	}

	submitDataBody := client.SubmitSelfServiceRegistrationFlowBody{
		SubmitSelfServiceRegistrationFlowWithPasswordMethodBody: client.NewSubmitSelfServiceRegistrationFlowWithPasswordMethodBody("password", pass, trait),
	}

	submitDataBody.SubmitSelfServiceRegistrationFlowWithPasswordMethodBody.SetCsrfToken(csrfToken)

	fmt.Println("SubmitRegistrationFlow", cookie, "test", csrfToken)
	apiClient := client.NewAPIClient(configuration)
	resp, r, err := apiClient.V0alpha2Api.SubmitSelfServiceRegistrationFlow(context.Background()).Flow(flowID).SubmitSelfServiceRegistrationFlowBody(submitDataBody).Cookie(cookie).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.SubmitSelfServiceRegistrationFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitSelfServiceRegistrationFlow`: SuccessfulSelfServiceRegistrationWithoutBrowser
	fmt.Fprintf(os.Stdout, "Response from `V0alpha2Api.SubmitSelfServiceRegistrationFlow`: %v\n", resp)
	return nil
}
