package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sdslabs/nymeria/pkg/workflow/registration"
)

func HandleGetRegistrationFlow(c *gin.Context) {
	cookie, flowID, csrf_token, err := registration.InitializeRegistrationFlowWrapper()

	if err != nil {
		fmt.Println(err)
		return
	}
	t := &registration.Traits{
		Email: "rohith@gmail.com",
	}
	t.Name.First = "Rohith"
	t.Name.Last = "Varma"

	err = registration.SubmitRegistrationFlowWrapper(cookie, flowID, csrf_token, *t)

	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "register",
	})

}
