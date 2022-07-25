package registeration

import (
	"github.com/sdslabs/nymeria/pkg/requests"
)

func InitializeRegisterationFlowWrapper() (string, string, error) {
	respBody := new(InitializeRegistration)
	setCookie, err := requests.GetInitializeFlowJSON("http://localhost:4433/self-service/registration/browser", respBody)

	if err != nil {
		return "", "", err
	}

}
