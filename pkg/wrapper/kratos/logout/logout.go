package logout

import (
	"context"
	"net/http"

	client "github.com/ory/client-go"

	"github.com/sdslabs/nymeria/config"
)

func InitializeLogoutFlowWrapper(cookie string) (*client.SelfServiceLogoutUrl, error) {
	apiClient := client.NewAPIClient(config.KratosClientConfig)

	resp, _, err := apiClient.V0alpha2Api.CreateSelfServiceLogoutFlowUrlForBrowsers(context.Background()).Cookie(cookie).Execute()

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func SubmitLogoutFlowWrapper(cookie string, token string, returnToUrl string) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:4433/self-service/logout", nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("token", token)

	req.URL.RawQuery = q.Encode()
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Accept", "application/json")

	_, err = client.Do(req)

	if err != nil {
		return err
	}

	return nil
}
