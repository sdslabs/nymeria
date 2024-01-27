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
		"email":         data.Email,
		"name":          data.Name,
		"password":      data.Password,
		"phone_number":  data.PhoneNumber,
		"img_url":       data.ImgURL,
		"invite_status": "pending",
		"verified":      false,
		"role":          data.Role,
		"created_at":    timeStamp,
		"totp_enabled":  false,
	}

	adminCreateIdentityBody := *client.NewAdminCreateIdentityBody("default", trait) // AdminCreateIdentityBody |  (optional)

	apiClient := client.NewAPIClient(config.KratosClientConfigAdmin)
	createdIdentity, r, err := apiClient.V0alpha2Api.AdminCreateIdentity(context.Background()).AdminCreateIdentityBody(adminCreateIdentityBody).Execute()

	return createdIdentity, r, err
}

func GetIdentityFlowWrapper(createdIdentity string) (*client.Identity, *http.Response, error) {
	apiClient := client.NewAPIClient(config.KratosClientConfigAdmin)

	getIdentity, r, err := apiClient.V0alpha2Api.AdminGetIdentity(context.Background(), createdIdentity).Execute()

	return getIdentity, r, err
}

func DeleteIdentityFlowWrapper(identity string) (*http.Response, error) {
	apiClient := client.NewAPIClient(config.KratosClientConfigAdmin)

	r, err := apiClient.V0alpha2Api.AdminDeleteIdentity(context.Background(), identity).Execute()

	return r, err
}

func ListIdentityFlowWrapper() ([]client.Identity, *http.Response, error) {
	apiClient := client.NewAPIClient(config.KratosClientConfigAdmin)

	identities, r, err := apiClient.V0alpha2Api.AdminListIdentities(context.Background()).Execute()

	return identities, r, err

}

func BanIdentityFlowWrapper(identity *client.Identity) (*client.Identity, *http.Response, error) {
	newState, err := client.NewIdentityStateFromValue("inactive")
	if err != nil {
		return nil, nil, err
	}

	submitDataBody := *client.NewAdminUpdateIdentityBody(identity.SchemaId, *newState, identity.Traits.(map[string]interface{}))

	apiClient := client.NewAPIClient(config.KratosClientConfigAdmin)
	id, r, err := apiClient.V0alpha2Api.AdminUpdateIdentity(context.Background(), identity.Id).AdminUpdateIdentityBody(submitDataBody).Execute()

	return id, r, err
}

func RemoveBanIdentityFlowWrapper(identity *client.Identity) (*client.Identity, *http.Response, error) {
	newState, err := client.NewIdentityStateFromValue("active")
	if err != nil {
		return nil, nil, err
	}

	submitDataBody := *client.NewAdminUpdateIdentityBody(identity.SchemaId, *newState, identity.Traits.(map[string]interface{}))

	apiClient := client.NewAPIClient(config.KratosClientConfigAdmin)
	id, r, err := apiClient.V0alpha2Api.AdminUpdateIdentity(context.Background(), identity.Id).AdminUpdateIdentityBody(submitDataBody).Execute()

	return id, r, err
}
