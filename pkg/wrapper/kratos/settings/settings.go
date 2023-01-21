package settings

import (
	"context"

	client "github.com/ory/client-go"

	"github.com/sdslabs/nymeria/config"
)

func InitializeSettingsFlowWrapper(req_cookie string) (client.SelfServiceSettingsFlow, string, error) {

	returnTo := "http://127.0.0.1:4455/ping" // string | The URL to return the browser to after the flow was completed. (optional)

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	resp, httpRes, err := apiClient.V0alpha2Api.InitializeSelfServiceSettingsFlowForBrowsers(context.Background()).ReturnTo(returnTo).Cookie(req_cookie).Execute()
	if err != nil {
		return *resp, "", err
	}

	cookie := httpRes.Header.Get("Host")

	return *resp, cookie, nil
}

func SubmitSettingsFlowWrapper(cookie string, flowID string, csrfToken string, TOTPcode string, TOTPUnlink bool) (string, string, error) {
	submitDataBody := client.SubmitSelfServiceSettingsFlowBody{SubmitSelfServiceSettingsFlowWithTotpMethodBody: client.NewSubmitSelfServiceSettingsFlowWithTotpMethodBody("totp")} // SubmitSelfServiceLoginFlowBody |

	submitDataBody.SubmitSelfServiceSettingsFlowWithTotpMethodBody.SetCsrfToken(csrfToken)
	submitDataBody.SubmitSelfServiceSettingsFlowWithTotpMethodBody.SetTotpCode(TOTPcode)
	submitDataBody.SubmitSelfServiceSettingsFlowWithTotpMethodBody.SetTotpUnlink(TOTPUnlink)

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	resp, r, err := apiClient.V0alpha2Api.SubmitSelfServiceSettingsFlow(context.Background()).Flow(flowID).SubmitSelfServiceSettingsFlowBody(submitDataBody).XSessionToken("").Cookie(cookie).Execute()
	if err != nil {
		return "", "", err
	}

	responseCookies := r.Header["Set-Cookie"]

	return resp.Id, responseCookies[1], nil
}
