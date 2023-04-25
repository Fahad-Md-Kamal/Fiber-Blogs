package configs

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
)

type EnvConfig struct {
	ServingPort  string `mapstructure:"SERVING_PORT"`
	DbHost       string `mapstructure:"DB_HOST"`
	DbPort       string `mapstructure:"DB_PORT"`
	DbName       string `mapstructure:"DB_NAME"`
	DbUser       string `mapstructure:"DB_USER"`
	DbPassword   string `mapstructure:"DB_PASSWORD"`
	SecretKey    string `mapstructure:"SECRET_KEY"`
	JwtSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

var ENVs EnvConfig

func LoadEnvs() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	envVars := make(map[string]string)
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		envVars[pair[0]] = pair[1]
	}

	// var cfg EnvConfig
	err = mapstructure.Decode(envVars, &ENVs)
	if err != nil {
		return fmt.Errorf("error decoding env vars: %w", err)
	}

	return nil
}
