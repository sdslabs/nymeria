package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/wrapper/kratos/registration"
)

func HandleGetRegistrationFlow(c *gin.Context) {
	cookie, flowID, csrf_token, err := registration.InitializeRegistrationFlowWrapper()

	if err != nil {
		log.ErrorLogger("Kratos get registration flow failed", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Kratos get registration flow failed",
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
		log.ErrorLogger("Unable to process json body", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Unable to process json body",
		})
		return
	}

	cookie, err := c.Cookie("registration_flow")

	if err != nil {
		log.ErrorLogger("Cookie not found", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "cookie not found",
		})
		return
	}

	session, err := registration.SubmitRegistrationFlowWrapper(cookie, t.FlowID, t.CsrfToken, t.Password, t.Traits)

	if err != nil {
		log.ErrorLogger("Kratos post registration flow failed", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Kratos post registration flow failed",
		})
		return
	}

	c.SetCookie("sdslabs_session", session, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"status": "created",
	})

}
