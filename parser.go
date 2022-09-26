package main

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

type conf struct {
	Url struct{
    Frontend_url string `yaml:"frontend_url"`
    Backend_url string `yaml:"backend_url"`
	} `yaml: "url"`
}

func (c *conf) getConf() *conf {

    yamlFile, err := ioutil.ReadFile("config.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
		log.Fatalf("Unmarshal: %v", err)
    }

    return c
}

func main() {
    var c conf
    c.getConf()

    fmt.Println(c.Url.Frontend_url)
	fmt.Println(c.Url.Backend_url)
}