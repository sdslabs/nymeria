package controller

import(
	"context"
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
	//m"github.com/sdslabs/nymeria/pkg/middleware"
	client "github.com/ory/kratos-client-go"
)

func CreateIdentity(c *gin.Context) {
    configuration := client.NewConfiguration()
    configuration.Servers = []client.ServerConfiguration{
        {
            URL: "http://127.0.0.1:4434", // Kratos Admin API
        },
    }
    apiClient := client.NewAPIClient(configuration)
    adminCreateIdentityBody := *client.NewAdminCreateIdentityBody(
        "default",
        map[string]interface{}{
            "email": "foo2@example.com",
            "name": map[string]string{
                "first": "foo2",
                "last":  "bar2",
            },
        },
    ) // AdminCreateIdentityBody |  (optional)

    createdIdentity, r, err := apiClient.V0alpha2Api.AdminCreateIdentity(context.Background()).AdminCreateIdentityBody(adminCreateIdentityBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.AdminCreateIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminCreateIdentity`: Identity
    fmt.Fprintf(os.Stdout, "Created identity with ID: %v\n", createdIdentity.Id)

}

func GetIdentity(c *gin.Context){
    configuration := client.NewConfiguration()
    configuration.Servers = []client.ServerConfiguration{
        {
            URL: "http://127.0.0.1:4434", // Kratos Admin API
        },
    }
    apiClient := client.NewAPIClient(configuration)
    createdIdentity := "32ff6997-04b0-46e4-a368-aa6d415bc410"
	getIdentity, r, err := apiClient.V0alpha2Api.AdminGetIdentity(context.Background(), createdIdentity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.AdminGetIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    fmt.Fprintf(os.Stdout, "Email for identity with id %v. Traits %v\n", createdIdentity, getIdentity.Traits)
}

func DeleteIdentity(c *gin.Context){
    configuration := client.NewConfiguration()
    configuration.Servers = []client.ServerConfiguration{
        {
            URL: "http://127.0.0.1:4434", // Kratos Admin API
        },
    }
    apiClient := client.NewAPIClient(configuration)

    identity := "32ff6997-04b0-46e4-a368-aa6d415bc410"

	r, err := apiClient.V0alpha2Api.AdminDeleteIdentity(context.Background(), identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.AdminDeleteIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    fmt.Println("Successfully Removed identity")
}