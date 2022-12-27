package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	client "github.com/ory/client-go"
	"github.com/sdslabs/nymeria/log"
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
		log.ErrorLogger("Error while calling `AdminCreateIdentity`", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "INternal server error",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"identity": createdIdentity.Id,
	})

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
		log.ErrorLogger("Error while calling `AdminGetIdentity`", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "INternal server error",
		})
	}

	jsonString, _ := json.Marshal(getIdentity.Traits)

	var identity Identity
	if err := json.Unmarshal(jsonString, &identity); err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(os.Stdout, "Identity details for id %v. Traits: %v\n", createdIdentity, identity)
	c.JSON(http.StatusOK, gin.H{
		"Identity": createdIdentity,
		"Traits": identity,
	})
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
		log.ErrorLogger("Error while calling `AdminDeleteIdentity`", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "INternal server error",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "removed identity",
	})
}

func ListIdentity(c *gin.Context) {
	configuration := client.NewConfiguration()
    configuration.Servers = []client.ServerConfiguration{
        {
            URL: "http://127.0.0.1:4434", // Kratos Admin API
        },
    }
    apiClient := client.NewAPIClient(configuration)

	identities, r, err := apiClient.V0alpha2Api.AdminListIdentities(context.Background()).Execute()

	if err != nil {
		log.ErrorLogger("Error while calling `AdminListIdentities`", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		c.JSON(http.StatusInternalServerError, gin.H {
			"error" : "Internal server error",
		})
	}
	c.JSON(http.StatusOK, gin.H {
		"identities": identities,
	})
}

func UpdateBanIdentity(c *gin.Context) {
	configuration := client.NewConfiguration()
    configuration.Servers = []client.ServerConfiguration{
        {
            URL: "http://127.0.0.1:4434", // Kratos Admin API
        },
    }
    apiClient := client.NewAPIClient(configuration)

	identity := c.PostForm("identity")

	jsonPatch := []client.JsonPatch{
		{	
			From: nil,
			Op: "replace",
			Path: "/active",
			Value: false,
		},
	}
	id, r, err := apiClient.V0alpha2Api.AdminPatchIdentity(context.Background(), identity).JsonPatch(jsonPatch).Execute()

	if err != nil {
		log.ErrorLogger("Error while calling `AdminPatchIdentities`", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		c.JSON(http.StatusInternalServerError, gin.H {
			"error" : err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H {
		"identities": id,
	})
}