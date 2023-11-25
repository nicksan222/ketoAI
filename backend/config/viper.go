package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvVars struct {
	AUTH0_DOMAIN   string `mapstructure:"AUTH0_DOMAIN"`
	AUTH0_AUDIENCE string `mapstructure:"AUTH0_AUDIENCE"`
	PORT           string `mapstructure:"PORT"`
	OPENAI_API_KEY string `mapstructure:"OPENAI_API_KEY"`
}

func LoadConfig() (config EnvVars, err error) {
	godotenv.Load()

	return EnvVars{
		AUTH0_DOMAIN:   os.Getenv("AUTH0_DOMAIN"),
		AUTH0_AUDIENCE: os.Getenv("AUTH0_AUDIENCE"),
		PORT:           os.Getenv("PORT"),
		OPENAI_API_KEY: os.Getenv("OPENAI_API_KEY"),
	}, err
}
