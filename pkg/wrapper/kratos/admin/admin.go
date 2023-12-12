package admin

import (
	"context"
	"net/http"

	"github.com/ory/client-go"

	"github.com/sdslabs/nymeria/config"
)

func CreateIdentityFlowWrapper(identityMap map[string]interface{}) (*client.Identity, *http.Response, error) {
	apiClient := client.NewAPIClient(config.KratosClientConfigAdmin)

	adminCreateIdentityBody := *client.NewAdminCreateIdentityBody(
		"default",
		identityMap,
	) // AdminCreateIdentityBody |  (optional)

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

func BanIdentityFlowWrapper(identity string) (*client.Identity, *http.Response, error) {
	apiClient := client.NewAPIClient(config.KratosClientConfigAdmin)

	jsonPatch := []client.JsonPatch{
		{
			From:  nil,
			Op:    "replace",
			Path:  "/active",
			Value: false,
		},
	}
	id, r, err := apiClient.V0alpha2Api.AdminPatchIdentity(context.Background(), identity).JsonPatch(jsonPatch).Execute()

	return id, r, err
}
