package main

// import (
// 	"github.com/sdslabs/nymeria/api"
// )

// func main() {
// 	api.Start()
// }

import (
	"github.com/sdslabs/nymeria/pkg/workflow/login"
)

func main() {
	login.InitializeLoginFlowWrapper()
}
