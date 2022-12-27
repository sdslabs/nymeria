package oidc

import (
	"context"
	"fmt"
	"os"

	client "github.com/ory/kratos-client-go"
	"github.com/sdslabs/nymeria/config"
)

func SubmitOIDCRegistrationFlowWrapper(provider string, cookie string, flowID string, csrfToken string) (string, error) {
	submitOIDCRegistrationFlowBody := client.SubmitSelfServiceRegistrationFlowBody{
		SubmitSelfServiceRegistrationFlowWithOidcMethodBody: client.NewSubmitSelfServiceRegistrationFlowWithOidcMethodBody("oidc", provider),
	}

	submitOIDCRegistrationFlowBody.SubmitSelfServiceRegistrationFlowWithOidcMethodBody.SetCsrfToken(csrfToken)

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	_, r, err := apiClient.V0alpha2Api.SubmitSelfServiceRegistrationFlow(context.Background()).Flow(flowID).SubmitSelfServiceRegistrationFlowBody(submitOIDCRegistrationFlowBody).Cookie(cookie).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.SubmitSelfServiceRegistrationFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseCookies := r.Header["Set-Cookie"]
	return responseCookies[1], nil
}