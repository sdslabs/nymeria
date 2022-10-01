package logout

import (
    "context"
    "fmt"
    "os"
    client "github.com/ory/kratos-client-go"
)

func InitializeLogoutFlowWrapper (cookie string) (*client.SelfServiceLogoutUrl, error) {
    configuration := client.NewConfiguration()
    
	apiClient := client.NewAPIClient(configuration)
    
	resp, r, err := apiClient.V0alpha2Api.CreateSelfServiceLogoutFlowUrlForBrowsers(context.Background()).Cookie(cookie).Execute()
    
	if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.CreateSelfServiceLogoutFlowUrlForBrowsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return resp, err
    }
    
	// response from `CreateSelfServiceLogoutFlowUrlForBrowsers`: SelfServiceLogoutUrl
    fmt.Fprintf(os.Stdout, "Response from `V0alpha2Api.CreateSelfServiceLogoutFlowUrlForBrowsers`: %v\n", resp)

	return resp, nil
}

func SubmitLogoutFlowWrapper (logoutUrl *client.SelfServiceLogoutUrl, returnToUrl string) {
	token := logoutUrl.LogoutToken
	
	configuration := client.NewConfiguration()
    
	apiClient := client.NewAPIClient(configuration)
    
	resp, err := apiClient.V0alpha2Api.SubmitSelfServiceLogoutFlow(context.Background()).Token(token).ReturnTo(returnToUrl).Execute()
    
	if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.SubmitSelfServiceLogoutFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
