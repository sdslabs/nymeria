package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sdslabs/nymeria/pkg/workflow/login"
)

func HandleGetLoginFlow(c *gin.Context) {
	cookie, flowID, csrf_token, err := login.InitializeLoginFlowWrapper()

	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
	}

	cookie, err := c.Cookie("login_flow")

	if err != nil {
		fmt.Println(err)
	}

	err = login.SubmitLoginFlowWrapper(cookie, t.FlowID, t.CsrfToken, t.Password, t.Identifier)

	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"flowID": "test",
	})

}
