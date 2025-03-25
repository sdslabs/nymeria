package login

import (
	"context"

	client "github.com/ory/client-go"

	"github.com/sdslabs/nymeria/config"
	"github.com/sdslabs/nymeria/helper"
)

func InitializeLoginFlowWrapper(aal string, cookie string) (string, string, string, error) {
	refresh := false // bool | Refresh a login session  If set to true, this will refresh an existing login session by asking the user to sign in again. This will reset the authenticated_at time of the session. (optional)
	returnTo := ""   // string | The URL to return the browser to after the flow was completed. (optional)

	apiClient := client.NewAPIClient(config.KratosClientConfig)

	resp, r, err := apiClient.FrontendAPI.CreateBrowserLoginFlow(context.Background()).Refresh(refresh).Aal(aal).ReturnTo(returnTo).Cookie(cookie).Execute()

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

func SubmitLoginFlowWrapper(cookie string, flowID string, csrfToken string, pass string, identifier string) (client.Session, string, string, error) {

	submitDataBody := client.UpdateLoginFlowBody{UpdateLoginFlowWithPasswordMethod: client.NewUpdateLoginFlowWithPasswordMethod(identifier, "password", pass)} // SubmitSelfServiceLoginFlowBody |
	submitDataBody.UpdateLoginFlowWithPasswordMethod.SetCsrfToken(csrfToken)

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	resp, r, err := apiClient.FrontendAPI.UpdateLoginFlow(context.Background()).Cookie(cookie).Flow(flowID).XSessionToken("").UpdateLoginFlowBody(submitDataBody).Execute()

	responseCookies := r.Header["Set-Cookie"]

	if err != nil {
		msg := helper.ExtractErrorMessage(r)
		if responseCookies == nil {
			return *client.NewSessionWithDefaults(), "", msg, err
		}
		return *client.NewSessionWithDefaults(), responseCookies[1], msg, err
	}

	return resp.Session, responseCookies[1], "", nil
}

func SubmitLoginWithMFAWrapper(cookie string, flowID string, csrfToken string, totp string) (client.Session, string, string, error) {
	submitDataBody := client.UpdateLoginFlowBody{UpdateLoginFlowWithTotpMethod: client.NewUpdateLoginFlowWithTotpMethod("totp", totp)} // SubmitSelfServiceLoginFlowBody |

	submitDataBody.UpdateLoginFlowWithTotpMethod.SetCsrfToken(csrfToken)

	apiClient := client.NewAPIClient(config.KratosClientConfig)

	resp, r, err := apiClient.FrontendAPI.UpdateLoginFlow(context.Background()).Flow(flowID).UpdateLoginFlowBody(submitDataBody).XSessionToken("").Cookie(cookie).Execute()

	errMsg := helper.ExtractErrorMessage(r)

	if err != nil {
		return *client.NewSessionWithDefaults(), "", errMsg, err
	}

	responseCookies := r.Header["Set-Cookie"]

	return resp.Session, responseCookies[0], "", nil
}
