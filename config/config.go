package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

type AppConfig struct {
	PORT         string
	POSTGRES_URI string
	GO_ENV       string
}

// func (ac AppConfig) String() string {
// 	jsonRep, _ := json.Marshal(ac)
// 	return string(jsonRep)
// }

func getenv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		log.Panicf("Could not find environment variable with key: %v", key)
	}

	return value
}

func InitAppConfig() AppConfig {
	env := os.Getenv("GO_ENV")

	if env == "" {
		env = "development"
	}
	projectDirName := "ratemyclass-be"
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + "/.env." + env)
	if err != nil {
		log.Printf("Could not load .env.%v file. Using os environments", env)
	}

	return AppConfig{PORT: getenv("PORT"), GO_ENV: env, POSTGRES_URI: getenv("POSTGRES_URI")}
}

var Module = fx.Module("config", fx.Provide(InitAppConfig))
