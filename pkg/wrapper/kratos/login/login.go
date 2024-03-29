package login

import (
	"context"

	client "github.com/ory/client-go"

	"github.com/sdslabs/nymeria/config"
)

func InitializeLoginFlowWrapper(aal string, cookie string) (string, string, string, error) {
	refresh := false                         // bool | Refresh a login session  If set to true, this will refresh an existing login session by asking the user to sign in again. This will reset the authenticated_at time of the session. (optional)
	returnTo := "http://localhost:4455/ping" // string | The URL to return the browser to after the flow was completed. (optional)

	apiClient := client.NewAPIClient(config.KratosClientConfig)

	resp, r, err := apiClient.V0alpha2Api.InitializeSelfServiceLoginFlowForBrowsers(context.Background()).Refresh(refresh).Aal(aal).ReturnTo(returnTo).Cookie(cookie).Execute()

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

func SubmitLoginFlowWrapper(cookie string, flowID string, csrfToken string, pass string, identifier string) (client.Session, string, error) {
	submitDataBody := client.SubmitSelfServiceLoginFlowBody{SubmitSelfServiceLoginFlowWithPasswordMethodBody: client.NewSubmitSelfServiceLoginFlowWithPasswordMethodBody(identifier, "password", pass)} // SubmitSelfServiceLoginFlowBody |

	submitDataBody.SubmitSelfServiceLoginFlowWithPasswordMethodBody.SetCsrfToken(csrfToken)

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	resp, r, err := apiClient.V0alpha2Api.SubmitSelfServiceLoginFlow(context.Background()).Flow(flowID).SubmitSelfServiceLoginFlowBody(submitDataBody).XSessionToken("").Cookie(cookie).Execute()
	if err != nil {
		return *client.NewSessionWithDefaults(), "", err
	}

	responseCookies := r.Header["Set-Cookie"]

	return resp.Session, responseCookies[1], nil
}

func SubmitLoginWithMFAWrapper(cookie string, flowID string, csrfToken string, totp string) (client.Session, string, error) {
	submitDataBody := client.SubmitSelfServiceLoginFlowBody{SubmitSelfServiceLoginFlowWithTotpMethodBody: client.NewSubmitSelfServiceLoginFlowWithTotpMethodBody("totp", totp)} // SubmitSelfServiceLoginFlowBody |

	submitDataBody.SubmitSelfServiceLoginFlowWithTotpMethodBody.SetCsrfToken(csrfToken)

	apiClient := client.NewAPIClient(config.KratosClientConfig)

	resp, r, err := apiClient.V0alpha2Api.SubmitSelfServiceLoginFlow(context.Background()).Flow(flowID).SubmitSelfServiceLoginFlowBody(submitDataBody).XSessionToken("").Cookie(cookie).Execute()

	if err != nil {
		return *client.NewSessionWithDefaults(), "", err
	}

	responseCookies := r.Header["Set-Cookie"]

	return resp.Session, responseCookies[0], nil
}
