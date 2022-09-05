package registration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sdslabs/nymeria/pkg/requests"
)

func InitializeRegistrationFlowWrapper() (string, string, string, error) {
	respBody := new(InitializeRegistration)
	setCookie, err := requests.GetInitializeFlowJSON("http://localhost:4433/self-service/registration/browser", respBody)

	if err != nil {
		return "", "", "", err
	}

	var csrf_token string

	for _, node := range respBody.UI.Nodes {
		if node.Attributes.Name == "csrf_token" {
			csrf_token = node.Attributes.Value
			break
		}
	}

	return setCookie, respBody.ID, csrf_token, nil
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
