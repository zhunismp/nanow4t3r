package config

type DBConfig struct {
	HOST     string
	PORT     string
	USERNAME string
	PASSWORD string
	NAME     string
}

type AppConfig struct {
	PORT string
}

type Config struct {
	APP_CONFIG AppConfig
	DB_CONFIG  DBConfig
}
