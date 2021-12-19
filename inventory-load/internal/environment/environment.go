package environment

import (
	"context"
	"fmt"
	"os"
)

type Environment struct {
	DbUser    string
	DbPwd     string
	DbTCPHost string
	DbPort    string
	DbName    string
}

func GetEnvironment(ctx context.Context) (*Environment, error) {
	env := &Environment{
		DbUser:    getEnv("DB_USER", "postgres"),
		DbPwd:     getEnv("DB_PASSWORD", "mysecretpassword"),
		DbTCPHost: getEnv("DB_HOST", "127.0.0.1"),
		DbPort:    getEnv("DB_PORT", "5432"),
		DbName:    getEnv("DB_NAME", "postgres"),
	}

	return env, nil

}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		fmt.Println("value from ENV", key, value)
		return value
	}

	return defaultVal
}
