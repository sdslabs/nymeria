package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/wrapper/kratos/logout"
)

func HandleGetLogoutFlow(c *gin.Context) {
	cookie, err := c.Cookie("sdslabs_session")

	if err != nil {
		log.ErrorLogger("Session cookie not found", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Session cookie not found",
		})
		return
	}

	logoutUrl, err := logout.InitializeLogoutFlowWrapper(cookie)

	if err != nil {
		log.ErrorLogger("Kratos get logout flow failed", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Kratos get logout flow failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"logout_token": logoutUrl.LogoutToken,
		"logout_url":   logoutUrl.LogoutUrl,
	})
}

func HandlePostLogoutFlow(c *gin.Context) {
	var t logout.SubmitLogoutBody
	err := c.BindJSON(&t)

	if err != nil {
		log.ErrorLogger("Unable to process json body", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Unable to process json body",
		})
		return
	}

	cookie, err := c.Cookie("sdslabs_session")

	if err != nil {
		log.ErrorLogger("Session cookie not found", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Session cookie not found",
		})
		return
	}
	err = logout.SubmitLogoutFlowWrapper(cookie, t.LogoutToken, t.LogoutUrl)

	if err != nil {
		log.ErrorLogger("Kratos get logout flow failed", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Kratos get logout flow failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "user logged out",
	})
}
