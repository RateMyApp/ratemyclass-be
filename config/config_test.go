package config

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func Test_AppConfig_ShouldFailWhenNoEnvIsPresent(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Did not panic")
		}
	}()

	InitAppConfig()
}

func Test_AppConfig_ShouldPass_WhenAllEnvVarArePresentInDotEnvFile(t *testing.T) {
	expectedOutput := AppConfig{
		PORT:      "2000",
		MONGO_URI: "mongodb://localhost:27017",
		GO_ENV:    "testing",
	}
	defer func() {
		os.Remove(".env.testing")
	}()
	os.Setenv("GO_ENV", "testing")
	f, err := os.Create(".env.testing")
	if err != nil {
		t.Error("Could not create .env.testing file")
		f.Close()
	}

	envStr := ""

	reflectVal := reflect.ValueOf(expectedOutput)

	for i := 0; i < reflectVal.NumField(); i++ {
		envStr += fmt.Sprintf("%v=%v\n", reflectVal.Type().Field(i).Name, reflectVal.Field(i).Interface())
	}

	f.Write([]byte(envStr))
	f.Close()

	appConfig := InitAppConfig()

	isEqual := reflect.DeepEqual(appConfig, expectedOutput)

	if !isEqual {
		t.Errorf("AppConfig Error: expected %v but got %v", expectedOutput, appConfig)
	}
}
