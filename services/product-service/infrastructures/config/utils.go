package config

// TODO: realize this function to load config from environment variables or a config file
func LoadConfig() Config {
	appConfig := AppConfig{
		PORT: "8080",
	}

	dbConfig := DBConfig{
		HOST:     "localhost",
		PORT:     "5432",
		USERNAME: "root",
		PASSWORD: "password",
		NAME:     "product_db",
	}

	return Config{
		APP_CONFIG: appConfig,
		DB_CONFIG:  dbConfig,
	}
}
