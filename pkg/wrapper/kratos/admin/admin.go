package admin

import (
	"context"
	"net/http"

	client "github.com/ory/client-go"

	"github.com/sdslabs/nymeria/config"
	"github.com/sdslabs/nymeria/pkg/middleware"
)

func CreateIdentityFlowWrapper(data Identity) (*client.Identity, *http.Response, error) {
	timeStamp := middleware.CurrentTimeStamp()

	trait := map[string]interface{}{
		"username":      data.Username,
		"email":         data.Email,
		"name":          data.Name,
		"phone_number":  data.PhoneNumber,
		"img_url":       data.ImgURL,
		"github_id":     data.GithubID,
		"invited_by":    data.InvitedBy,
		"invite_status": "pending",
		"role":          data.Role,
		"created_at":    timeStamp,
		"totp_enabled":  false,
	}

	adminCreateIdentityBody := *client.NewCreateIdentityBody("default", trait) // AdminCreateIdentityBody |  (optional)

	apiClient := client.NewAPIClient(config.KratosClientConfigAdmin)
	createdIdentity, r, err := apiClient.IdentityAPI.CreateIdentity(context.Background()).CreateIdentityBody(adminCreateIdentityBody).Execute()

	return createdIdentity, r, err
}

func GetIdentityFlowWrapper(createdIdentity string) (*client.Identity, *http.Response, error) {
	apiClient := client.NewAPIClient(config.KratosClientConfigAdmin)

	getIdentity, r, err := apiClient.IdentityAPI.GetIdentity(context.Background(), createdIdentity).Execute()

	return getIdentity, r, err
}

func DeleteIdentityFlowWrapper(identity string) (*http.Response, error) {
	apiClient := client.NewAPIClient(config.KratosClientConfigAdmin)

	r, err := apiClient.IdentityAPI.DeleteIdentity(context.Background(), identity).Execute()

	return r, err
}

func ListIdentityFlowWrapper() ([]client.Identity, *http.Response, error) {
	apiClient := client.NewAPIClient(config.KratosClientConfigAdmin)

	identities, r, err := apiClient.IdentityAPI.ListIdentities(context.Background()).Execute()

	return identities, r, err

}

func BanIdentityFlowWrapper(identity *client.Identity) (*client.Identity, *http.Response, error) {

	submitDataBody := *client.NewUpdateIdentityBody(identity.SchemaId, "inactive", identity.Traits.(map[string]interface{}))

	apiClient := client.NewAPIClient(config.KratosClientConfigAdmin)
	id, r, err := apiClient.IdentityAPI.UpdateIdentity(context.Background(), identity.Id).UpdateIdentityBody(submitDataBody).Execute()

	return id, r, err
}

func RemoveBanIdentityFlowWrapper(identity *client.Identity) (*client.Identity, *http.Response, error) {

	submitDataBody := *client.NewUpdateIdentityBody(identity.SchemaId, "active", identity.Traits.(map[string]interface{}))

	apiClient := client.NewAPIClient(config.KratosClientConfigAdmin)
	id, r, err := apiClient.IdentityAPI.UpdateIdentity(context.Background(), identity.Id).UpdateIdentityBody(submitDataBody).Execute()

	return id, r, err
}

func RoleSwitchFlowWrapper(identity *client.Identity) (*client.Identity, *http.Response, error) {
	traits := identity.GetTraits().(map[string]interface{})

	if traits["role"] == "user" {
		traits["role"] = "admin"
	} else if traits["role"] == "admin" {
		traits["role"] = "user"
	}

	submitDataBody := *client.NewUpdateIdentityBody(identity.SchemaId, *identity.State, traits)

	apiClient := client.NewAPIClient(config.KratosClientConfigAdmin)
	id, r, err := apiClient.IdentityAPI.UpdateIdentity(context.Background(), identity.Id).UpdateIdentityBody(submitDataBody).Execute()

	return id, r, err
}
