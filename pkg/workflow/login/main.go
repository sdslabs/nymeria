package login

import (
	"context"
	"fmt"
	"os"

	client "github.com/ory/kratos-client-go"
)

func InitializeLoginFlowWrapper() (string, string, string, error) {
	refresh := false                         // bool | Refresh a login session  If set to true, this will refresh an existing login session by asking the user to sign in again. This will reset the authenticated_at time of the session. (optional)
	aal := "aal1"                            // string | Request a Specific AuthenticationMethod Assurance Level  Use this parameter to upgrade an existing session's authenticator assurance level (AAL). This allows you to ask for multi-factor authentication. When an identity sign in using e.g. username+password, the AAL is 1. If you wish to \"upgrade\" the session's security by asking the user to perform TOTP / WebAuth/ ... you would set this to \"aal2\". (optional)
	returnTo := "http://127.0.0.1:4455/ping" // string | The URL to return the browser to after the flow was completed. (optional)

	configuration := client.NewConfiguration()
	configuration.Servers = []client.ServerConfiguration{
		{
			URL: "http://127.0.0.1:4433",
		},
	}
	apiClient := client.NewAPIClient(configuration)
	resp, fullr, err := apiClient.V0alpha2Api.InitializeSelfServiceLoginFlowForBrowsers(context.Background()).Refresh(refresh).Aal(aal).ReturnTo(returnTo).Execute()
	cookie := fullr.Header.Get("Set-Cookie")

	if err != nil {
		return "", "", "", err
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

func SubmitLoginFlowWrapper(cookie string, flowID string, csrfToken string, pass string, identifier string) (string, error) {
	submitDataBody := client.SubmitSelfServiceLoginFlowBody{SubmitSelfServiceLoginFlowWithPasswordMethodBody: client.NewSubmitSelfServiceLoginFlowWithPasswordMethodBody(identifier, "password", pass)} // SubmitSelfServiceLoginFlowBody |

	csrf_token := csrfToken
	submitDataBody.SubmitSelfServiceLoginFlowWithPasswordMethodBody.SetCsrfToken(csrf_token)

	configuration := client.NewConfiguration()
	configuration.Servers = []client.ServerConfiguration{
		{
			URL: "http://127.0.0.1:4433",
		},
	}
	apiClient := client.NewAPIClient(configuration)
	_, r, err := apiClient.V0alpha2Api.SubmitSelfServiceLoginFlow(context.Background()).Flow(flowID).SubmitSelfServiceLoginFlowBody(submitDataBody).XSessionToken("").Cookie(cookie).Execute()
	if err != nil {
		fmt.Println(r)
		return "", err
	}

	responseCookies := r.Header["Set-Cookie"]

	return responseCookies[1], nil
}
