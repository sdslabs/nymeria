package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sdslabs/nymeria/pkg/controller/logout"
	"go.uber.org/zap"
)

func HandleGetLogoutFlow(c *gin.Context) {
	cookie, err := c.Cookie("sdslabs_session")

	if err != nil {
		zap.L().Error("Session cookie not found", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "session not found",
		})
		return
	}

	logoutUrl, err := logout.InitializeLogoutFlowWrapper(cookie)

	if err != nil {
		zap.L().Error("Kratos get logout flow failed", zap.Error(err))
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
		zap.L().Error("Unable to process json body", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to process request body",
		})
		return
	}

	cookie, err := c.Cookie("sdslabs_session")

	if err != nil {
		zap.L().Error("Session cookie not found", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "session not found",
		})
		return
	}
	err = logout.SubmitLogoutFlowWrapper(cookie, t.LogoutToken, t.LogoutUrl)

	if err != nil {
		zap.L().Error("Kratos get logout flow failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "user logged out",
	})
}
