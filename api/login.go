package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sdslabs/nymeria/pkg/workflow/login"
)

func HandleGetLoginFLow(c *gin.Context) {
	cookie, flowID, csrf_token, err := login.InitializeLoginFlowWrapper()

	if err != nil {
		fmt.Println(err)
		return
	}
	t := &login.Traits{
		Email:    "rohith@gmail.com",
		Password: "jngkjenrjg",
	}

	err = login.SubmitLoginFlowWrapper(cookie, flowID, csrf_token, *t)

	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login",
	})

}
