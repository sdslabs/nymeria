package config

type Config struct {
	URL URL `yaml:"url"`
}

type URL struct {
	FrontendURL string `yaml:"frontend_url"`
	BackendURL  string `yaml:"kratos_url"`
}
