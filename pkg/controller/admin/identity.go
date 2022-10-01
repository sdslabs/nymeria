package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	client "github.com/ory/kratos-client-go"
	m "github.com/sdslabs/nymeria/pkg/middleware"
)

func CreateIdentity(c *gin.Context) {
	apiClient := m.NewAdminMiddleware()
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
	apiClient := m.NewAdminMiddleware()
	createdIdentity := "80b8317c-a1be-4510-b415-987f28b7667b"
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
	apiClient := m.NewAdminMiddleware()

	identity := "5e2d9d8c-8367-478b-b183-268bd4a88bf1"

	r, err := apiClient.V0alpha2Api.AdminDeleteIdentity(context.Background(), identity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.AdminDeleteIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	fmt.Println("Successfully Removed identity")
}

func GetAllIdentity(c *gin.Context) {

}
