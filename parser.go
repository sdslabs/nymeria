package main

import (
	"fmt"
	"io/ioutil"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type config struct {
	url struct  {
		frontend_url string `yaml:"frontend_url"`;
		backend_url string `yaml:"backend_url"`;
	} `yaml:"url"`
}

func main() {
	vi := viper.New()
	vi.SetConfigFile("config.yaml")
	vi.ReadInConfig()
	fmt.Println(vi.GetString("url.frontend_url"))
	
	C := &config{}
	source,err1 := ioutil.ReadFile("config.yaml")
	if err1 != nil {
		fmt.Println(err1)
	}
	
	err := yaml.Unmarshal([]byte(source),&C)
	if err != nil {
		fmt.Println("unable to decode into struct, %v", err)
	}

	fmt.Println(C.url.backend_url)

}