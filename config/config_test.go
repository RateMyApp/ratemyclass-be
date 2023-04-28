package config_test

import (
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/ratemyapp/config"
	"github.com/stretchr/testify/assert"
)

func Test_AppConfig_ShouldPanic_WhenNoEnvIsPresent(t *testing.T) {
	// log.Println("hello")
	os.Setenv("GO_ENV", "asdfasdfahfdkja")

	reflectElem := reflect.ValueOf(config.AppConfig{})

	fieldsNo := reflectElem.NumField()

	for i := 0; i < fieldsNo; i++ {
		field := reflectElem.Type().Field(i).Name
		log.Println(field)
		if field == "GO_ENV" {
			continue
		}
		os.Setenv(field, "")
	}

	defer func() {
		for i := 0; i < fieldsNo; i++ {
			field := reflectElem.Type().Field(i).Name
			os.Unsetenv(field)
		}
		if r := recover(); r == nil {
			t.Errorf("Did not panic")
		}
	}()
	appConfig := config.InitAppConfig()
	log.Println(appConfig.POSTGRES_URI)
}

func Test_AppConfig_ShouldPass_WhenAllEnvVarArePresentInDotEnvFile(t *testing.T) {
	// expectedOutput := AppConfig{
	// 	PORT:         "2000",
	// 	POSTGRES_URI: "postgres://user:pass@localhost:5432/db_name",
	// 	GO_ENV:       "testing",
	// }
	// defer func() {
	// 	os.Remove(".env.testing")
	// }()
	os.Setenv("GO_ENV", "testing")
	// f, err := os.Create(".env.testing")
	// if err != nil {
	// 	t.Error("Could not create .env.testing file")
	// 	f.Close()
	// }
	//
	// envStr := ""
	//
	// reflectVal := reflect.ValueOf(expectedOutput)
	//
	// for i := 0; i < reflectVal.NumField(); i++ {
	// 	envStr += fmt.Sprintf("%v=%v\n", reflectVal.Type().Field(i).Name, reflectVal.Field(i).Interface())
	// }
	//
	// f.Write([]byte(envStr))
	// f.Close()

	appConfig := config.InitAppConfig()

	assert := assert.New(t)

	assert.NotEmpty(appConfig.GO_ENV, "AppConfig Error: Expected field GO_ENV to not be empty")
	assert.NotEmpty(appConfig.POSTGRES_URI, "AppConfig Error: Expected field POSTGRES_URI to not be empty")
	assert.NotEmpty(appConfig.PORT, "AppConfig Error: Expected field PORT to not be empty")
	// isEqual := reflect.DeepEqual(appConfig, expectedOutput)

	// if !isEqual {
	// 	t.Errorf("AppConfig Error: expected %v but got %v", expectedOutput.String(), appConfig.String())
	// }
}
