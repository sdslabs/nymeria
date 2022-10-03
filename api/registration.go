package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sdslabs/nymeria/pkg/controller/registration"
	"go.uber.org/zap"
)

func HandleGetRegistrationFlow(c *gin.Context) {
	cookie, flowID, csrf_token, err := registration.InitializeRegistrationFlowWrapper()

	if err != nil {
		zap.L().Error("Kratos get registration flow failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.SetCookie("registration_flow", cookie, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"flowID":     flowID,
		"csrf_token": csrf_token,
	})

}

func HandlePostRegistrationFlow(c *gin.Context) {
	var t registration.SubmitRegistrationBody
	err := c.BindJSON(&t)

	if err != nil {
		zap.L().Error("Unable to process json body", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to process request body",
		})
		return
	}

	cookie, err := c.Cookie("registration_flow")

	if err != nil {
		zap.L().Error("Cookie not found", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "csrf cookie not found",
		})
		return
	}

	session, err := registration.SubmitRegistrationFlowWrapper(cookie, t.FlowID, t.CsrfToken, t.Password, t.Traits)

	if err != nil {
		zap.L().Error("Kratos post registration flow failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.SetCookie("sdslabs_session", session, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"status": "created",
	})

}
