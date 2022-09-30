package config

type Config struct {
	Url Url `yaml: "url"`
}

type Url struct {
	Frontend_url string `yaml:"frontend_url"`
	Backend_url  string `yaml:"kratos_url"`
}
