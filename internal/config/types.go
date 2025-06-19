package config

// DatabaseConfig holds database configuration
type Config struct {
	DBHost     string `env:"DB_HOST"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`
	DBPort     string `env:"DB_PORT"`

	CSRFSecret string `env:"CSRF_SECRET"`
	CSRFMaxAge int    `env:"CSRF_MAX_AGE"` // in minutes
	JWTSecret  string `env:"JWT_SECRET"`
	JWTMaxAge  int    `env:"JWT_MAX_AGE"` // in days

	EnvMode string `env:"ENV_MODE"`

	MailFrom     string `env:"MAIL_FROM"`
	MailPassword string `env:"MAIL_PASSWORD"`
	SMTPHost     string `env:"SMTP_HOST"`
	SMTPPort     string `env:"SMTP_PORT"`
	OTPExpiry    int    `env:"OTP_EXPIRY"` // in minutes
}
