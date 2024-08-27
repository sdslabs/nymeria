package settings

import (
	"context"
	"fmt"
	"os"
	"strings"

	client "github.com/ory/client-go"

	"github.com/sdslabs/nymeria/config"
)

func InitializeSettingsFlowWrapper(session_cookie string) (client.SettingsFlow, string, error) {

	returnTo := "" // string | The URL to return the browser to after the flow was completed. (optional)

	cookie := strings.Split(session_cookie, ";")[0]

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	resp, httpRes, err := apiClient.FrontendAPI.CreateBrowserSettingsFlow(context.Background()).ReturnTo(returnTo).Cookie(cookie).Execute()

	if err != nil {
		return *client.NewSettingsFlowWithDefaults(), "", err
	}

	cookie = httpRes.Header.Get("Set-Cookie")

	return *resp, cookie, nil
}

func SubmitSettingsFlowPasswordMethod(flow_cookie string, session_cookie string, flowID string, csrfToken string, password string) (string, error) {
	submitFlowBody := client.UpdateSettingsFlowBody{
		UpdateSettingsFlowWithPasswordMethod: client.NewUpdateSettingsFlowWithPasswordMethod("password", password),
	}

	submitFlowBody.UpdateSettingsFlowWithPasswordMethod.SetCsrfToken(csrfToken)
	cookie := strings.Split(flow_cookie, ";")[0] + "; " + strings.Split(session_cookie, ";")[0] + "; x-csrf-token=" + csrfToken

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	_, r, err := apiClient.FrontendAPI.UpdateSettingsFlow(context.Background()).Flow(flowID).Cookie(cookie).UpdateSettingsFlowBody(submitFlowBody).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.SubmitSelfServiceVerificationFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return "", err
	}

	return "Password Changed", nil
}

func SubmitSettingsFlowProfileMethod(flow_cookie string, session_cookie string, flowID string, csrfToken string, traits map[string]interface{}) (string, error) {
	submitFlowBody := client.UpdateSettingsFlowBody{
		UpdateSettingsFlowWithProfileMethod: client.NewUpdateSettingsFlowWithProfileMethod("profile", traits),
	}

	submitFlowBody.UpdateSettingsFlowWithProfileMethod.SetCsrfToken(csrfToken)

	cookie := strings.Split(flow_cookie, ";")[0] + "; " + strings.Split(session_cookie, ";")[0] + "; x-csrf-token=" + csrfToken

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	_, r, err := apiClient.FrontendAPI.UpdateSettingsFlow(context.Background()).Flow(flowID).Cookie(cookie).UpdateSettingsFlowBody(submitFlowBody).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.SubmitSelfServiceVerificationFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return "", err
	}

	return "Profile Updated", nil
}

func SubmitSettingsFlowTOTPMethod(flow_cookie string, session_cookie string, flowID string, csrfToken string, TOTPcode string, TOTPUnlink bool) (string, error) {
	submitFlowBody := client.UpdateSettingsFlowBody{
		UpdateSettingsFlowWithTotpMethod: client.NewUpdateSettingsFlowWithTotpMethod("totp"),
	}

	submitFlowBody.UpdateSettingsFlowWithTotpMethod.SetCsrfToken(csrfToken)
	submitFlowBody.UpdateSettingsFlowWithTotpMethod.SetTotpCode(TOTPcode)
	submitFlowBody.UpdateSettingsFlowWithTotpMethod.SetTotpUnlink(TOTPUnlink)

	cookie := strings.Split(flow_cookie, ";")[0] + "; " + strings.Split(session_cookie, ";")[0] + "; x-csrf-token=" + csrfToken

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	_, r, err := apiClient.FrontendAPI.UpdateSettingsFlow(context.Background()).Flow(flowID).Cookie(cookie).UpdateSettingsFlowBody(submitFlowBody).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.SubmitSelfServiceVerificationFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return "", err
	}

	return "Totp Toggled", nil
}
