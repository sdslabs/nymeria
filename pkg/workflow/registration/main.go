package registration

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	client "github.com/ory/kratos-client-go"
)

func InitializeRegistrationFlowWrapper() (string, string, string, error) {
	returnTo := "http://127.0.0.1:4455/ping"

	configuration := client.NewConfiguration()
	configuration.Servers = []client.ServerConfiguration{
		{
			URL: "http://127.0.0.1:4433",
		},
	}

	apiClient := client.NewAPIClient(configuration)
	resp, r, err := apiClient.V0alpha2Api.InitializeSelfServiceRegistrationFlowForBrowsers(context.Background()).ReturnTo(returnTo).Execute()
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

func SubmitRegistrationFlowWrapper(cookie string, flowID string, csrfToken string, data Traits) error {
	requestBody := new(SubmitRegistrationBody)
	requestBody.Method = "password"
	requestBody.Password = "jngkjenrjg"
	requestBody.CsrfToken = csrfToken
	requestBody.Data = data

	jsonData, err := json.Marshal(requestBody)

	if err != nil {
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:4433/self-service/registration", bytes.NewReader(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Cookie", cookie)
	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("flow", flowID)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	fmt.Println("Hello")
	if err != nil {
		return err
	}

	fmt.Println(resp)
	return nil
}
