package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sdslabs/nymeria/pkg/controller/login"
	"go.uber.org/zap"
)

func HandleGetLoginFlow(c *gin.Context) {
	cookie, flowID, csrf_token, err := login.InitializeLoginFlowWrapper()

	if err != nil {
		zap.L().Error("Kratos get login flow failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.SetCookie("login_flow", cookie, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"flowID":     flowID,
		"csrf_token": csrf_token,
	})
}

func HandlePostLoginFlow(c *gin.Context) {
	var t login.SubmitLoginAPIBody
	err := c.BindJSON(&t)

	if err != nil {
		zap.L().Error("Unable to process json body", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to process request body",
		})
		return
	}

	cookie, err := c.Cookie("login_flow")

	if err != nil {
		zap.L().Error("Cookie not found", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "csrf cookie not found",
		})
		return
	}

	session, err := login.SubmitLoginFlowWrapper(cookie, t.FlowID, t.CsrfToken, t.Password, t.Identifier)

	if err != nil {
		zap.L().Error("Kratos post login flow failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.SetCookie("sdslabs_session", session, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"status": "user logged in",
	})

}
