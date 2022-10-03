package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sdslabs/nymeria/pkg/controller/logout"
)

func HandleGetLogoutFlow(c *gin.Context) {
	cookie, err := c.Cookie("sdslabs_session")

	if err != nil {
		fmt.Println(err)

	}

	logoutUrl, err := logout.InitializeLogoutFlowWrapper(cookie)

	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
	}

	cookie, err := c.Cookie("sdslabs_session")

	if err != nil {
		fmt.Println(err)

	}
	err = logout.SubmitLogoutFlowWrapper(cookie, t.LogoutToken, t.LogoutUrl)

	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": "dorime",
	})
}
