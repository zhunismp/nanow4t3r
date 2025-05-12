package config

import (
	"log"
	"os"
)

func LoadConfig() Config {
	appConfig := AppConfig{
		PORT: getEnv("PRODUCT_APP_PORT", "8080"),
	}

	dbConfig := DBConfig{
		HOST:     getEnv("PRODUCT_DB_HOST", "product-db"),
		PORT:     getEnv("PRODUCT_DB_PORT", "5432"),
		USERNAME: getEnv("PRODUCT_DB_USER", "root"),
		PASSWORD: getEnv("PRODUCT_DB_PASSWORD", "password"),
		NAME:     getEnv("PRODUCT_DB_NAME", "product_db"),
	}

	return Config{
		APP_CONFIG: appConfig,
		DB_CONFIG:  dbConfig,
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	log.Default().Println("Using fallback for", key, ": ", fallback)
	return fallback
}
