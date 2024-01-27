package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	client "github.com/ory/client-go"

	"github.com/sdslabs/nymeria/config"
	"github.com/sdslabs/nymeria/log"
)

func GetSession(c *gin.Context) (*client.Session, error) {
	cookie, err := c.Cookie("sdslabs_session")
	if err != nil {
		log.ErrorLogger("Cookie not found", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"message": "cookie not found",
		})
		return nil, err
	}
	apiClient := client.NewAPIClient(config.KratosClientConfig)
	resp, r, err := apiClient.V0alpha2Api.ToSession(context.Background()).Cookie(cookie).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.ToSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return nil, err
	}
	return resp, nil
}

func OnlyAdmin(c *gin.Context) {
	session, err := GetSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	identity := session.GetIdentity()
	traits := identity.GetTraits()
	role := traits.(map[string]interface{})["role"]
	if role == "admin" {
		c.Next()
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	c.Abort()
}

func OnlyUser(c *gin.Context) {
	session, err := GetSession(c)
	if err != nil {
		c.Abort()
		return
	}
	identity := session.GetIdentity()
	traits := identity.GetTraits()
	role := traits.(map[string]interface{})["role"]
	if role == "user" {
		c.Next()
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	c.Abort()
}

func SuperAdmin(c *gin.Context) {
	session, err := GetSession(c)
	if err != nil {
		c.Abort()
		return
	}
	identity := session.GetIdentity()
	traits := identity.GetTraits()
	role := traits.(map[string]interface{})["role"]
	if role == "superadmin" {
		c.Next()
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	c.Abort()
}
