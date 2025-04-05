package registration

import (
	"context"

	client "github.com/ory/client-go"

	"github.com/sdslabs/nymeria/config"
	"github.com/sdslabs/nymeria/helper"
	"github.com/sdslabs/nymeria/pkg/middleware"
)

func InitializeRegistrationFlowWrapper() (string, string, string, error) {
	returnTo := ""

	apiClient := client.NewAPIClient(config.KratosClientConfig)

	resp, r, err := apiClient.FrontendAPI.CreateBrowserRegistrationFlow(context.Background()).ReturnTo(returnTo).Execute()
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

func SubmitRegistrationFlowWrapper(cookie string, flowID string, csrfToken string, password string, data Traits) (string, []string, string, error) {
	timeStamp := middleware.CurrentTimeStamp()
	trait := map[string]interface{}{
		"username":      data.Username,
		"email":         data.Email,
		"name":          data.Name,
		"img_url":       data.ImgURL,
		"phone_number":  data.PhoneNumber,
		"github_id":     data.GithubID,
		"invited_by":    data.InvitedBy,
		"invite_status": "self_created",
		"role":          "user",
		"created_at":    timeStamp,
		"totp_enabled":  false,
	}

	submitDataBody := client.UpdateRegistrationFlowBody{UpdateRegistrationFlowWithPasswordMethod: client.NewUpdateRegistrationFlowWithPasswordMethod("password", password, trait)}

	submitDataBody.UpdateRegistrationFlowWithPasswordMethod.SetCsrfToken(csrfToken)

	apiClient := client.NewAPIClient(config.KratosClientConfig)

	resp, r, err := apiClient.FrontendAPI.UpdateRegistrationFlow(context.Background()).Flow(flowID).UpdateRegistrationFlowBody(submitDataBody).Cookie(cookie).Execute()

	if err != nil {
		msg := helper.ExtractErrorMessage(r)

		return "", nil, msg, err
	}

	responseCookies := r.Header["Set-Cookie"]

	return resp.GetContinueWith()[1].ContinueWithVerificationUi.GetFlow().Id, responseCookies, "", nil
}
