package configs

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

const (
	prod = "production"
)

// Config object
type Config struct {
	Env  string `env:"ENV"`
	Host string `env:"APP_HOST"`
	Port string `env:"APP_PORT"`
}

// IsProd Checks if env is production
func (c Config) IsProd() bool {
	return c.Env == prod
}

// LoadConfig gets config from .env
func LoadConfig() {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	environmentPath := filepath.Join(currentPath, ".env")

	if err := godotenv.Load(environmentPath); err != nil {
		log.Fatal("Error loading .env file")
		log.Fatal(err)
	}
}

// GetConfig gets all config for the application
func GetConfig() Config {
	return Config{
		Env:  os.Getenv("ENV"),
		Host: os.Getenv("APP_HOST"),
		Port: os.Getenv("APP_PORT"),
	}
}
