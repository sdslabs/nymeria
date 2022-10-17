package login

import (
	"context"

	client "github.com/ory/kratos-client-go"
	"github.com/sdslabs/nymeria/config"
)

func InitializeLoginFlowWrapper() (string, string, string, error) {
	refresh := false                         // bool | Refresh a login session  If set to true, this will refresh an existing login session by asking the user to sign in again. This will reset the authenticated_at time of the session. (optional)
	aal := "aal1"                            // string | Request a Specific AuthenticationMethod Assurance Level  Use this parameter to upgrade an existing session's authenticator assurance level (AAL). This allows you to ask for multi-factor authentication. When an identity sign in using e.g. username+password, the AAL is 1. If you wish to \"upgrade\" the session's security by asking the user to perform TOTP / WebAuth/ ... you would set this to \"aal2\". (optional)
	returnTo := "http://127.0.0.1:4455/ping" // string | The URL to return the browser to after the flow was completed. (optional)

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	resp, httpRes, err := apiClient.V0alpha2Api.InitializeSelfServiceLoginFlowForBrowsers(context.Background()).Refresh(refresh).Aal(aal).ReturnTo(returnTo).Execute()
	if err != nil {
		return "", "", "", err
	}

	cookie := httpRes.Header.Get("Host")

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

	submitDataBody.SubmitSelfServiceLoginFlowWithPasswordMethodBody.SetCsrfToken(csrfToken)

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	_, r, err := apiClient.V0alpha2Api.SubmitSelfServiceLoginFlow(context.Background()).Flow(flowID).SubmitSelfServiceLoginFlowBody(submitDataBody).XSessionToken("").Cookie(cookie).Execute()
	if err != nil {
		return "", err
	}

	responseCookies := r.Header["Set-Cookie"]

	return responseCookies[1], nil
}
