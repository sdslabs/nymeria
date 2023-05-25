package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/sdslabs/nymeria/config"
	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/wrapper/kratos/registration"
)

func HandleGetRegistrationFlow(c *gin.Context) {
	cookie, flowID, csrf_token, err := registration.InitializeRegistrationFlowWrapper()

	if err != nil {
		log.ErrorLogger("Kratos get registration flow failed", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Kratos get registration flow failed",
		})
		return
	}

	c.SetCookie("registration_flow", cookie, 3600, "/", config.NymeriaConfig.URL.Domain, true, true)

	c.JSON(http.StatusOK, gin.H{
		"flowID":     flowID,
		"csrf_token": csrf_token,
	})

}

func HandlePostRegistrationFlow(c *gin.Context) {
	var t registration.SubmitRegistrationBody
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

	cookie, err := c.Cookie("registration_flow")

	if err != nil {
		log.ErrorLogger("Cookie not found", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "cookie not found",
		})
		return
	}

	session, err := registration.SubmitRegistrationFlowWrapper(cookie, t.FlowID, t.CsrfToken, t.Password, t.Traits)

	if err != nil {
		log.ErrorLogger("Kratos post registration flow failed", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Kratos post registration flow failed",
		})
		return
	}

	c.SetCookie("sdslabs_session", session, 3600, "/", config.NymeriaConfig.URL.Domain, true, true)
	c.JSON(http.StatusOK, gin.H{
		"status": "created",
	})

}
