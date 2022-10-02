package main

import (
	"fmt"

	"github.com/sdslabs/nymeria/pkg/workflow/login"
)

// import (
// 	"github.com/sdslabs/nymeria/api"
// )

// func main() {
// 	api.Start()
// }

func main() {
	a, b, c, d := login.InitializeLoginFlowWrapper()

	fmt.Println(a, b, c)
	if d != nil {
		fmt.Print(d)
	}

	fmt.Println("second req")
	data := login.Traits{
		Email: "test@test.com",
	}
	data.Name.First = "pratham"
	data.Name.Last = "ks"
	login.SubmitLoginFlowWrapper(a, b, c, "pass", data)
}
