package main

// import (
// 	"github.com/sdslabs/nymeria/api"
// )

// func main() {
// 	api.Start()
// }

import (
	"fmt"

	"github.com/sdslabs/nymeria/pkg/workflow/registration"
)

func main() {
	a, b, c, err := registration.InitializeRegistrationFlowWrapper()

	fmt.Println("Initial", a, b, c)
	if err != nil {
		fmt.Print("Uhuihiu")
	}
}
