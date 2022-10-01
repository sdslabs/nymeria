package login

import (
	"context"
	"fmt"
	"os"

	client "github.com/ory/kratos-client-go"
)

func InitializeLoginFlowWrapper() (string, string, string, error) {
	refresh := true                     // bool | Refresh a login session  If set to true, this will refresh an existing login session by asking the user to sign in again. This will reset the authenticated_at time of the session. (optional)
	aal := "aal1"                       // string | Request a Specific AuthenticationMethod Assurance Level  Use this parameter to upgrade an existing session's authenticator assurance level (AAL). This allows you to ask for multi-factor authentication. When an identity sign in using e.g. username+password, the AAL is 1. If you wish to \"upgrade\" the session's security by asking the user to perform TOTP / WebAuth/ ... you would set this to \"aal2\". (optional)
	returnTo := "http://localhost:4455" // string | The URL to return the browser to after the flow was completed. (optional)

	configuration := client.NewConfiguration()
	configuration.Servers = []client.ServerConfiguration{
		{
			URL: "http://127.0.0.1:4455",
		},
	}
	apiClient := client.NewAPIClient(configuration)
	resp, fullr, err := apiClient.V0alpha2Api.InitializeSelfServiceLoginFlowForBrowsers(context.Background()).Refresh(refresh).Aal(aal).ReturnTo(returnTo).Execute()
	cookie := fullr.Header.Get("Set-Cookie")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.InitializeSelfServiceLoginFlowForBrowsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", fullr)
	}
	// response from `InitializeSelfServiceLoginFlowForBrowsers`: SelfServiceLoginFlow
	fmt.Fprintf(os.Stdout, "Response from `V0alpha2Api.InitializeSelfServiceLoginFlowForBrowsers`: %v\n", resp)

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

func SubmitLoginFlowWrapper(cookie string, flowID string, csrfToken string, data Traits) error {
	flow := flowID                                                                                                                                                                                                                          // string | The Login Flow ID  The value for this parameter comes from `flow` URL Query parameter sent to your application (e.g. `/login?flow=abcde`).
	submitSelfServiceLoginFlowBody := client.SubmitSelfServiceLoginFlowBody{SubmitSelfServiceLoginFlowWithLookupSecretMethodBody: client.NewSubmitSelfServiceLoginFlowWithLookupSecretMethodBody("LookupSecret_example", "Method_example")} // SubmitSelfServiceLoginFlowBody |
	xSessionToken := csrfToken
	cookies := cookie // string | The Session Token of the Identity performing the settings flow. (optional)                                                                                                                                                                                                                      // string | HTTP Cookies  When using the SDK in a browser app, on the server side you must include the HTTP Cookie Header sent by the client to your server here. This ensures that CSRF and session cookies are respected. (optional)

	configuration := client.NewConfiguration()
	configuration.Servers = []client.ServerConfiguration{
		{
			URL: "http://127.0.0.1:4455",
		},
	}
	apiClient := client.NewAPIClient(configuration)
	resp, r, err := apiClient.V0alpha2Api.SubmitSelfServiceLoginFlow(context.Background()).Flow(flow).SubmitSelfServiceLoginFlowBody(submitSelfServiceLoginFlowBody).XSessionToken(xSessionToken).Cookie(cookies).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.SubmitSelfServiceLoginFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitSelfServiceLoginFlow`: SuccessfulSelfServiceLoginWithoutBrowser
	fmt.Fprintf(os.Stdout, "Response from `V0alpha2Api.SubmitSelfServiceLoginFlow`: %v\n", resp)
	return nil
}
