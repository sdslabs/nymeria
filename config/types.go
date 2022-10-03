package config

type NymeriaCfg struct {
	URL URL `yaml:"url"`
}

type URL struct {
	FrontendURL string `yaml:"frontend_url"`
	KratosURL   string `yaml:"kratos_url"`
}
