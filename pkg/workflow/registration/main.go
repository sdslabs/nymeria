package registration

import (
	"github.com/sdslabs/nymeria/pkg/requests"
)

func InitializeRegisterationFlowWrapper() (string, string, string, error) {
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

func SubmitRegistrationFlowWrapper(cookie string, flowID string, csrf_token string, data Traits) {

}
