package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/db"
	"github.com/sdslabs/nymeria/pkg/middleware"
)

type Profile struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

func HandlePostProfile(c *gin.Context) {
	var body AccessProfileRequest
	err := c.BindJSON(&body)
	if err != nil {
		log.ErrorLogger("Unable to process json body", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Unable to process json body",
		})
		return
	}
	app, err := db.GetApplication(body.ClientKey, body.ClientSecret)
	if err != nil {
		log.ErrorLogger("Unable to get application", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Internal Server Error",
		})
		return
	}
	if app.RedirectURL != body.RedirectURL {
		log.ErrorLogger("Redirect URL does not match", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Redirect URL does not match",
		})
		return
	}
	if app.ClientKey != body.ClientKey {
		log.ErrorLogger("Client Key does not match", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Unauthorized",
		})
		return
	}
	if app.ClientSecret != body.ClientSecret {
		log.ErrorLogger("Client Secret does not match", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Unauthorized",
		})
		return
	}
	session, err := middleware.GetSession(c)
	if err != nil {
		log.ErrorLogger("Unable to get session", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Unable to get session",
		})
		return
	}
	identity := session.GetIdentity()
	traits := identity.GetTraits()
	profile := traits.(map[string]interface{})
	response := Profile{
		Email:       profile["email"].(string),
		Name:        profile["name"].(string),
		PhoneNumber: profile["phone_number"].(string),
	}
	c.JSON(http.StatusOK, response)
}
