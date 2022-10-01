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
	// t := &registration.Traits{
	// 	Email: "rohith@gmail.com",
	// }
	// t.Name.First = "Rohith"
	// t.Name.Last = "Varma"

	c.SetCookie("registration_flow", cookie, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"flowID":     flowID,
		"csrf_token": csrf_token,
	})

}

func HandlePostRegistrationFlow(c *gin.Context) {
	var t registration.SubmitRegistrationAPIBody
	err := c.BindJSON(&t)

	if err != nil {
		fmt.Println(err)
	}

	cookie, err := c.Cookie("registration_flow")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cookie)

	err = registration.SubmitRegistrationFlowWrapper(cookie, t.FlowID, t.CsrfToken, t.Password, t.Traits)

	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"flowID": "test",
	})

}
