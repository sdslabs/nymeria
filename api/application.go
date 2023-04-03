package api

import (
	"net/http"
	"strconv"
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

		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Unable to get application data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": app,
	})

}

func HandlePostApplication(c *gin.Context) {
	var body ApplicationPostBody
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

	err = db.CreateApplication(body.Name, body.RedirectURL, body.AllowedDomains, body.Organisation, helper.RandomString(10), helper.RandomString(30))

	if err != nil {
		log.ErrorLogger("Create application failed", err)

		errCode, _ := strconv.Atoi((strings.Split(err.Error(), " "))[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Create application failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "application created",
	})

}

func HandlePutApplication(c *gin.Context) {
	var body ApplicationPutBody
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

	err = db.UpdateApplication(body.ID, body.Name, body.RedirectURL, body.AllowedDomains, body.Organisation)

	if err != nil {
		log.ErrorLogger("Update application failed", err)

		errCode, _ := strconv.Atoi((strings.Split(err.Error(), " "))[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Update application failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "application updated",
	})

}

func HandleDeleteApplication(c *gin.Context) {
	var body ApplicationBody
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

	err = db.DeleteApplication(body.ID)

	if err != nil {
		log.ErrorLogger("Delete application failed", err)

		errCode, _ := strconv.Atoi((strings.Split(err.Error(), " "))[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Delete application failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "application deleted",
	})

}

func HandleUpdateClientSecret(c *gin.Context) {
	var body ApplicationBody
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

	err = db.UpdateClientSecret(body.ID)

	if err != nil {
		log.ErrorLogger("Client Secret update failed", err)

		errCode, _ := strconv.Atoi((strings.Split(err.Error(), " "))[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Client Secret update failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Client Secret updated sucessfully",
	})

}
