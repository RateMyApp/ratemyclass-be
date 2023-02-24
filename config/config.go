package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

type AppConfig struct {
	PORT         string
	POSTGRES_URI string
	GO_ENV       string
}

func (ac AppConfig) String() string {
	jsonRep, _ := json.Marshal(ac)
	return string(jsonRep)
}

func getenv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Panicf("Could not find environment variable with key: %v", key)
	}

	return value
}

func InitAppConfig() AppConfig {
	env := os.Getenv("GO_ENV")

	if env == "" {
		env = "development"
	}

	err := godotenv.Load(".env." + env)
	if err != nil {
		log.Println("Could not load .env file. Using os environments")
	}

	return AppConfig{PORT: getenv("PORT"), GO_ENV: env, POSTGRES_URI: getenv("POSTGRES_URI")}
}

var Module = fx.Module("config", fx.Provide(InitAppConfig))
