package config

type Config struct {
	Application ApplicationConfig
	Database    DatabaseConfig
}

type ApplicationConfig struct {
	Port int
	Host string
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}
