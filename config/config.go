package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

type AppConfig struct {
	Port       string
	MongoURI string
	GoEnv      string
}

func getenv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Could not find environment variable with key: %v", key)
	}

	return value
}

func InitAppConfig() AppConfig {
	err := godotenv.Load(".env.development",)

	if err != nil {
		log.Println("Could not load .env file. Using os environments")
	}

	return AppConfig{Port: getenv("PORT"), GoEnv: getenv("GO_ENV"), MongoURI: getenv("MONGO_URI")}
}

var Module = fx.Module("config", fx.Provide(InitAppConfig))
