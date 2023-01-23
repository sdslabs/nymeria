package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/wrapper/kratos/verification"
)

func HandleGetVerificationFlow(c *gin.Context) {
	log.Logger.Debug("Get Verification")

	auth_cookie, _ := c.Cookie("sdslabs_session")

	cookie, flowID, csrf_token, err := verification.InitializeVerificationFlowWrapper(auth_cookie)

	if err != nil {
		log.ErrorLogger("Intialize Verification Failed", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Intialize Verification Failed",
		})
		return
	}

	c.SetCookie("verification_flow", cookie, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"flowID":     flowID,
		"csrf_token": csrf_token,
	})
}

func HandlePostVerificationFlow(c *gin.Context) {
	var t verification.SubmitVerificationBody
	err := c.Bind(&t)

	if err != nil {
		log.ErrorLogger("Unable to process json body", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Unable to process json body",
		})
		return
	}

	cookie, err := c.Cookie("verification_flow")
	session, err := c.Cookie("sdslabs_session")

	fmt.Println(cookie)
	fmt.Println(session)

	if err != nil {
		log.ErrorLogger("Cookie not found", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Cookie not found",
		})
		return
	}

	_, err = verification.SubmitVerificationFlowWrapper(cookie, session, t.FlowID, t.CsrfToken, t.Email, t.Method)

	if err != nil {
		log.ErrorLogger("Post Verification flow failed", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Post Settings flow failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Account Verification Successful",
	})
}
