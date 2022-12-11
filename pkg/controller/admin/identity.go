package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	client "github.com/ory/client-go"
)

func CreateIdentity(c *gin.Context) {
	configuration := client.NewConfiguration()
    configuration.Servers = []client.ServerConfiguration{
        {
            URL: "http://127.0.0.1:4434", // Kratos Admin API
        },
    }
    apiClient := client.NewAPIClient(configuration)

	id, _ := strconv.Atoi(c.PostForm("id"))
	verified, _ := strconv.Atoi(c.PostForm("verified"))
	active, _ := strconv.ParseBool(c.PostForm("active"))
	adminCreateIdentityBody := *client.NewAdminCreateIdentityBody(
		"default",
		map[string]interface{}{
			"id":           id,
			"name":         c.PostForm("name"),
			"username":     c.PostForm("username"),
			"email":        c.PostForm("email"),
			"phone_number": c.PostForm("phone_number"),
			"password":     c.PostForm("password"),
			"img_url":      c.PostForm("img_url"),
			"active":       active,
			"verified":     verified,
			"role":         c.PostForm("role"),
			"created_at":   c.PostForm("created_at"),
			"github_id":    c.PostForm("github_id"),
			"dribble_id":   c.PostForm("dribble_id"),
		},
	) // AdminCreateIdentityBody |  (optional)

	createdIdentity, r, err := apiClient.V0alpha2Api.AdminCreateIdentity(context.Background()).AdminCreateIdentityBody(adminCreateIdentityBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.AdminCreateIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	fmt.Fprintf(os.Stdout, "Created identity with ID: %v\n", createdIdentity.Id)

}

func GetIdentity(c *gin.Context) {
	configuration := client.NewConfiguration()
    configuration.Servers = []client.ServerConfiguration{
        {
            URL: "http://127.0.0.1:4434", // Kratos Admin API
        },
    }
    apiClient := client.NewAPIClient(configuration)

	createdIdentity := c.Query("identity")
	fmt.Println(createdIdentity)
	getIdentity, r, err := apiClient.V0alpha2Api.AdminGetIdentity(context.Background(), createdIdentity).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.AdminGetIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	jsonString, _ := json.Marshal(getIdentity.Traits)

	var identity Identity
	if err := json.Unmarshal(jsonString, &identity); err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(os.Stdout, "Identity details for id %v. Traits: %v\n", createdIdentity, identity)
}

func DeleteIdentity(c *gin.Context) {
	configuration := client.NewConfiguration()
    configuration.Servers = []client.ServerConfiguration{
        {
            URL: "http://127.0.0.1:4434", // Kratos Admin API
        },
    }
    apiClient := client.NewAPIClient(configuration)

	identity := c.PostForm("identity")

	r, err := apiClient.V0alpha2Api.AdminDeleteIdentity(context.Background(), identity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.AdminDeleteIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	fmt.Println("Successfully Removed identity")
}
