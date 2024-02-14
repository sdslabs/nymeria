package settings

import (
	"context"
	"fmt"
	"os"
	"strings"

	client "github.com/ory/client-go"

	"github.com/sdslabs/nymeria/config"
)

func InitializeSettingsFlowWrapper(session_cookie string, recovery_cookie string) (client.SelfServiceSettingsFlow, string, error) {

	returnTo := "http://localhost:4455/ping" // string | The URL to return the browser to after the flow was completed. (optional)

	var cookie string

	if recovery_cookie != "" {
		cookie = "ory_kratos_session=" + recovery_cookie
	} else {
		cookie = strings.Split(session_cookie, ";")[0]
	}

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	resp, httpRes, err := apiClient.V0alpha2Api.InitializeSelfServiceSettingsFlowForBrowsers(context.Background()).ReturnTo(returnTo).Cookie(cookie).Execute()
	if err != nil {
		return *client.NewSelfServiceSettingsFlowWithDefaults(), "", err
	}

	cookie = httpRes.Header.Get("Set-Cookie")

	return *resp, cookie, nil
}

func SubmitSettingsFlowPasswordMethod(flow_cookie string, session_cookie string, flowID string, csrfToken string, password string) (string, error) {
	submitFlowBody := client.SubmitSelfServiceSettingsFlowBody{
		SubmitSelfServiceSettingsFlowWithPasswordMethodBody: client.NewSubmitSelfServiceSettingsFlowWithPasswordMethodBody("password", password),
	}

	submitFlowBody.SubmitSelfServiceSettingsFlowWithPasswordMethodBody.SetCsrfToken(csrfToken)
	cookie := strings.Split(flow_cookie, ";")[0] + "; " + strings.Split(session_cookie, ";")[0] + "; x-csrf-token=" + csrfToken

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	_, r, err := apiClient.V0alpha2Api.SubmitSelfServiceSettingsFlow(context.Background()).Flow(flowID).Cookie(cookie).SubmitSelfServiceSettingsFlowBody(submitFlowBody).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.SubmitSelfServiceVerificationFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return "", err
	}

	return "Password Changed", nil
}

func SubmitSettingsFlowProfileMethod(flow_cookie string, session_cookie string, flowID string, csrfToken string, traits map[string]interface{}) (string, error) {
	submitFlowBody := client.SubmitSelfServiceSettingsFlowBody{
		SubmitSelfServiceSettingsFlowWithProfileMethodBody: client.NewSubmitSelfServiceSettingsFlowWithProfileMethodBody("profile", traits),
	}

	submitFlowBody.SubmitSelfServiceSettingsFlowWithProfileMethodBody.SetCsrfToken(csrfToken)

	cookie := strings.Split(flow_cookie, ";")[0] + "; " + strings.Split(session_cookie, ";")[0] + "; x-csrf-token=" + csrfToken

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	_, r, err := apiClient.V0alpha2Api.SubmitSelfServiceSettingsFlow(context.Background()).Flow(flowID).Cookie(cookie).SubmitSelfServiceSettingsFlowBody(submitFlowBody).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.SubmitSelfServiceVerificationFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return "", err
	}

	return "Profile Updated", nil
}

func SubmitSettingsFlowTOTPMethod(flow_cookie string, session_cookie string, flowID string, csrfToken string, TOTPcode string, TOTPUnlink bool) (string, error) {
	submitFlowBody := client.SubmitSelfServiceSettingsFlowBody{
		SubmitSelfServiceSettingsFlowWithTotpMethodBody: client.NewSubmitSelfServiceSettingsFlowWithTotpMethodBody("totp"),
	}

	submitFlowBody.SubmitSelfServiceSettingsFlowWithTotpMethodBody.SetCsrfToken(csrfToken)
	submitFlowBody.SubmitSelfServiceSettingsFlowWithTotpMethodBody.SetTotpCode(TOTPcode)
	submitFlowBody.SubmitSelfServiceSettingsFlowWithTotpMethodBody.SetTotpUnlink(TOTPUnlink)

	cookie := strings.Split(flow_cookie, ";")[0] + "; " + strings.Split(session_cookie, ";")[0] + "; x-csrf-token=" + csrfToken

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	_, r, err := apiClient.V0alpha2Api.SubmitSelfServiceSettingsFlow(context.Background()).Flow(flowID).Cookie(cookie).SubmitSelfServiceSettingsFlowBody(submitFlowBody).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.SubmitSelfServiceVerificationFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return "", err
	}

	return "Totp Toggled", nil
}
