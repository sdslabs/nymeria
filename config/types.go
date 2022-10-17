package config

type NymeriaCfg struct {
	Env string `yaml:"env"`
	URL URL    `yaml:"url"`
}

type URL struct {
	FrontendURL string `yaml:"frontend_url"`
	KratosURL   string `yaml:"kratos_url"`
}
