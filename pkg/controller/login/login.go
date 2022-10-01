package controller

import (
    "context"
    "fmt"
    "os"
    "github.com/gin-gonic/gin"
    client "github.com/ory/kratos-client-go"
)

func InitializeSelfServiceLoginFlow(c *gin.Context) {
    refresh := true // bool | Refresh a login session  If set to true, this will refresh an existing login session by asking the user to sign in again. This will reset the authenticated_at time of the session. (optional)
    aal := "aal1" // string | Request a Specific AuthenticationMethod Assurance Level  Use this parameter to upgrade an existing session's authenticator assurance level (AAL). This allows you to ask for multi-factor authentication. When an identity sign in using e.g. username+password, the AAL is 1. If you wish to \"upgrade\" the session's security by asking the user to perform TOTP / WebAuth/ ... you would set this to \"aal2\". (optional)
    returnTo := "http://127.0.0.1:4434" // string | The URL to return the browser to after the flow was completed. (optional)

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)

    resp, r, err := apiClient.V0alpha2Api.InitializeSelfServiceLoginFlowForBrowsers(context.Background()).Refresh(refresh).Aal(aal).ReturnTo(returnTo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.InitializeSelfServiceLoginFlowForBrowsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitializeSelfServiceLoginFlowForBrowsers`: SelfServiceLoginFlow
    fmt.Fprintf(os.Stdout, "Response from `V0alpha2Api.InitializeSelfServiceLoginFlowForBrowsers`: %v\n", resp)
}