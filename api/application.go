package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/sdslabs/nymeria/helper"
	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/db"
)

func HandleGetApplication(c *gin.Context) {
	app, err := db.GetAllApplication()

	if err != nil {
		log.ErrorLogger("Unable to get application data", err)
		errCode := helper.ExtractErrorCode(err)
		errStr := helper.ExtractErrorString(err)
		c.JSON(errCode, gin.H{
			"error":   http.StatusText(errCode),
			"message": "Unable to get application data: " + errStr,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": app,
	})

}

func HandleCreateApplication(c *gin.Context) {
	var body ApplicationPostBody
	err := c.BindJSON(&body)

	if err != nil {
		errCode := helper.ExtractErrorCode(err)

		log.ErrorLogger("Couldn't create application.", err)
		c.JSON(errCode, gin.H{
			"error":   http.StatusText(errCode),
			"message": "Unable to process json body",
		})
		return
	}

	clientKey := helper.RandomString(10)
	clientSecret := helper.RandomString(30)

	err = db.CreateApplication(body.Name, body.RedirectURL, body.AllowedDomains, body.Organization, clientKey, clientSecret)

	if err != nil {
		log.ErrorLogger("Create application failed", err)

		errCode := helper.ExtractErrorCode(err)
		errStr := helper.ExtractErrorString(err)
		c.JSON(errCode, gin.H{
			"error":   http.StatusText(errCode),
			"message": "Create application failed: " + errStr,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "application created",
		"application_credentials": gin.H{
			"name":            body.Name,
			"client_key":      clientKey,
			"client_secret":   clientSecret,
			"redirect_url":    body.RedirectURL,
			"allowed_domains": body.AllowedDomains,
			"organization":    body.Organization,
		},
	})

}

func HandleUpdateApplication(c *gin.Context) {
	var body ApplicationPutBody
	err := c.BindJSON(&body)

	if err != nil {
		log.ErrorLogger("Unable to process json body", err)

		errCode := helper.ExtractErrorCode(err)
		c.JSON(errCode, gin.H{
			"error":   http.StatusText(errCode),
			"message": "Unable to process json body",
		})
		return
	}

	err = db.UpdateApplication(body.ID, body.Name, body.RedirectURL, body.AllowedDomains, body.Organization)

	if err != nil {
		log.ErrorLogger("Update application failed", err)

		errCode := helper.ExtractErrorCode(err)
		errStr := helper.ExtractErrorString(err)
		c.JSON(errCode, gin.H{
			"error":   http.StatusText(errCode),
			"message": "Update application failed: " + errStr,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "application updated",
	})

}

func HandleDeleteApplication(c *gin.Context) {
	var body ApplicationBody
	err := c.BindJSON(&body)

	if err != nil {
		log.ErrorLogger("Unable to process json body", err)

		errCode := helper.ExtractErrorCode(err)
		c.JSON(errCode, gin.H{
			"error":   http.StatusText(errCode),
			"message": "Unable to process json body",
		})
		return
	}

	err = db.DeleteApplication(body.ID)

	if err != nil {
		log.ErrorLogger("Delete application failed", err)

		errCode := helper.ExtractErrorCode(err)
		errStr := helper.ExtractErrorString(err)
		c.JSON(errCode, gin.H{
			"error":   http.StatusText(errCode),
			"message": "Delete application failed: " + errStr,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "application deleted",
	})

}

func HandleUpdateClientSecret(c *gin.Context) {
	var body ApplicationBody
	err := c.BindJSON(&body)

	if err != nil {
		log.ErrorLogger("Unable to process json body", err)

		errCode := helper.ExtractErrorCode(err)
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Unable to process json body",
		})
		return
	}

	newSecret, err := db.UpdateClientSecret(body.ID)

	if err != nil {
		log.ErrorLogger("Client Secret update failed", err)

		errCode := helper.ExtractErrorCode(err)
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Client Secret update failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Client Secret updated successfully",
		"newSecret": newSecret,
	})

}

func HandleUpdateClientKey(c *gin.Context) {
	var body ApplicationBody
	err := c.BindJSON(&body)

	if err != nil {
		log.ErrorLogger("Unable to process json body", err)

		errCode := helper.ExtractErrorCode(err)
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Unable to process json body",
		})
		return
	}

	newKey, err := db.UpdateClientKey(body.ID)

	if err != nil {
		log.ErrorLogger("Client Key update failed", err)

		errCode := helper.ExtractErrorCode(err)
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Client Key update failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Client Key updated successfully",
		"newClientKey": newKey,
	})
}
