package registration

import (
	"context"
	"fmt"
	"os"

	client "github.com/ory/client-go"

	"github.com/sdslabs/nymeria/config"
	"github.com/sdslabs/nymeria/pkg/middleware"
)

func InitializeRegistrationFlowWrapper() (string, string, string, error) {
	returnTo := "http://localhost:4455/ping"

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
	timeStamp := middleware.CurrentTimeStamp()
	trait := map[string]interface{}{
		"email":         data.Email,
		"name":          data.Name,
		"password":      password,
		"img_url":       data.ImgURL,
		"phone_number":  data.PhoneNumber,
		"invite_status": "self_created",
		"verified":      false,
		"role":          "user",
		"created_at":    timeStamp,
		"totp_enabled":  false,
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
		return "", err
	}

	responseCookies := r.Header["Set-Cookie"]
	return responseCookies[1], nil
}
