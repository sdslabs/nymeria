package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sdslabs/nymeria/internal/database/applications"
	"github.com/sdslabs/nymeria/internal/database/schema"
	"github.com/sdslabs/nymeria/internal/utils"
)

func HandleGetApplicationFlow(c *gin.Context) {
	csrfToken, err := utils.GenerateCSRFToken(c.GetHeader("X-User-ID")) // TODO: Get user ID from session
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to generate CSRF token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"message":    "create application flow",
		"csrf_token": csrfToken,
	})
}

func HandleFetchAllApplicationsFlow(c *gin.Context) {
	apps, err := applications.GetAllApplications()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "applications fetched",
		"data":    apps,
	})
}

func HandleFetchApplicationByIDFlow(c *gin.Context) {
	id := c.Param("id")

	app, err := applications.GetApplicationByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "application fetched",
		"data":    app,
	})
}

func HandleCreateApplicationFlow(c *gin.Context) {
	var req CreateApplicationRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	clientKey, err := utils.GenerateRandomString(16)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to generate client key",
		})
		return
	}

	clientSecret, err := utils.GenerateRandomString(32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to generate client secret",
		})
		return
	}

	app := schema.Application{
		Name:           req.Name,
		AllowedOrigins: req.AllowedOrigins,
		RedirectURIs:   req.RedirectURIs,
		ClientKey:      clientKey,
		ClientSecret:   clientSecret,
	}

	err = applications.CreateApplication(app)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "application created",
		"data":    app,
	})
}

func HandleDeleteApplicationFlow(c *gin.Context) {
	id := c.Param("id")

	err := applications.DeleteApplication(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "application deleted",
	})
}

func HandleUpdateApplicationFlow(c *gin.Context) {
	var req UpdateApplicationRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	app, err := applications.GetApplicationByID(req.ApplicationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if req.NewKeyFlag {
		clientKey, err := utils.GenerateRandomString(16)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
			return
		}

		clientSecret, err := utils.GenerateRandomString(32)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
			return
		}

		app.ClientKey = clientKey
		app.ClientSecret = clientSecret
	}

	if req.ApplicationURL != "" {
		app.ApplicationURL = req.ApplicationURL
	}

	if len(req.AllowedOrigins) > 0 {
		app.AllowedOrigins = req.AllowedOrigins
	}

	if len(req.RedirectURIs) > 0 {
		app.RedirectURIs = req.RedirectURIs
	}

	err = applications.UpdateApplication(req.ApplicationID, app)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "application updated",
		"data": gin.H{
			"application_id": req.ApplicationID,
			"client_key":     app.ClientKey,
			"client_secret":  app.ClientSecret,
		},
	})
}
