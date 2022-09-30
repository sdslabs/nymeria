package main

// import (
// 	"github.com/sdslabs/nymeria/api"
// )

// func main() {
// 	api.Start()
// }

import (
	"fmt"

	"github.com/sdslabs/nymeria/config"
)

func main() {
	var c config.Config
	c.GetConf()

	fmt.Println(c.Url.Frontend_url)
	fmt.Println(c.Url.Backend_url)
}
