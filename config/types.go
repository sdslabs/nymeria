package config

type NymeriaCfg struct {
	Env string `yaml:"env"`
	URL URL    `yaml:"url"`
	DB  DB     `yaml:"db"`
}
type URL struct {
	KratosURL      string `yaml:"kratos_url"`
	AdminKratosURL string `yaml:"admin_kratos_url"`
	Domain         string `yaml:"domain"`
}

type DB struct {
	DSN      string `yaml:"dsn"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
}
