package settings

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/goccy/go-json"
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

func SubmitSettingsFlowPasswordMethod(flow_cookie string, session_cookie string, flowID string, csrfToken string, method string, password string) (string, error) {
	client := &http.Client{}

	var req_body SubmitSettingsWithPasswordBody
	req_body.Method = method
	req_body.Password = password
	req_body.CsrfToken = csrfToken

	body, err := json.Marshal(req_body)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:4433/self-service/settings", bytes.NewBuffer(body))

	q := req.URL.Query()
	q.Add("flow", flowID)

	if err != nil {
		return "", err
	}

	cookie := strings.Split(flow_cookie, ";")[0] + "; " + strings.Split(session_cookie, ";")[0] + "; x-csrf-token=" + csrfToken
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 200 {
		error := errors.New(resp.Status)
		return "", error
	}

	return "Password Changed", nil
}

func SubmitSettingsFlowProfileMethod(flow_cookie string, session_cookie string, flowID string, csrfToken string, method string, traits map[string]interface{}) (string, error) {
	client := &http.Client{}

	var req_body SubmitSettingsProfileRequest
	req_body.Method = method
	req_body.CsrfToken = csrfToken
	req_body.Traits = traits

	body, err := json.Marshal(req_body)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:4433/self-service/settings", bytes.NewBuffer(body))
	q := req.URL.Query()
	q.Add("flow", flowID)

	if err != nil {
		return "", err
	}
	cookie := strings.Split(flow_cookie, ";")[0] + "; " + strings.Split(session_cookie, ";")[0] + "; x-csrf-token=" + csrfToken
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 200 {
		error := errors.New(resp.Status)
		return "", error
	}

	return "Profile Updated", nil
}

func SubmitSettingsFlowTOTPMethod(flow_cookie string, session_cookie string, flowID string, csrfToken string, method string, TOTPcode string, TOTPUnlink bool) (string, error) {
	client := &http.Client{}

	var req_body SubmitSettingsWithTOTPBody
	req_body.Method = method
	req_body.TotpCode = TOTPcode
	req_body.TotpUnlink = TOTPUnlink
	req_body.CsrfToken = csrfToken

	body, err := json.Marshal(req_body)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:4433/self-service/settings", bytes.NewBuffer(body))

	q := req.URL.Query()
	q.Add("flow", flowID)

	if err != nil {
		return "", err
	}

	cookie := strings.Split(flow_cookie, ";")[0] + "; " + strings.Split(session_cookie, ";")[0] + "; x-csrf-token=" + csrfToken
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 200 {
		error := errors.New(resp.Status)
		return "", error
	}

	return "Totp Toggled", nil
}
