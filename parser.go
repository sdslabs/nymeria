package main

import (
	"fmt"
	"github.com/spf13/viper"
  )

func main() {
	v :=viper.New()
	v.SetConfigFile("parser.yaml")
	v.ReadInConfig()
	fmt.Println(v.GetString("url.frontend_url"))
	fmt.Println(v.GetString("url.backend_url"))

}