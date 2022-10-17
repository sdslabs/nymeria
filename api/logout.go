package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/wrapper/kratos/logout"
)

func HandleGetLogoutFlow(c *gin.Context) {
	cookie, err := c.Cookie("sdslabs_session")

	if err != nil {
		log.ErrorLogger("Session cookie not found", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "session not found",
		})
		return
	}

	logoutUrl, err := logout.InitializeLogoutFlowWrapper(cookie)

	if err != nil {
		log.ErrorLogger("Kratos get logout flow failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"logoutToken": logoutUrl.LogoutToken,
		"url":         logoutUrl.LogoutUrl,
	})
}

func HandlePostLogoutFlow(c *gin.Context) {
	var t logout.SubmitLogoutBody
	err := c.BindJSON(&t)

	if err != nil {
		log.ErrorLogger("Unable to process json body", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to process request body",
		})
		return
	}

	cookie, err := c.Cookie("sdslabs_session")

	if err != nil {
		log.ErrorLogger("Session cookie not found", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "session not found",
		})
		return
	}
	err = logout.SubmitLogoutFlowWrapper(cookie, t.LogoutToken, t.LogoutUrl)

	if err != nil {
		log.ErrorLogger("Kratos get logout flow failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "user logged out",
	})
}
