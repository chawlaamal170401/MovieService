package config

type DatabaseConfig struct {
	Host     string
	Driver   string
	User     string
	Password string
	DBName   string
	Port     int
	SSLMode  string
}
