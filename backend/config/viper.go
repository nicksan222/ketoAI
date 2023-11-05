package config

import "os"

type EnvVars struct {
	AUTH0_DOMAIN   string `mapstructure:"AUTH0_DOMAIN"`
	AUTH0_AUDIENCE string `mapstructure:"AUTH0_AUDIENCE"`
	PORT           string `mapstructure:"PORT"`
}

func LoadConfig() (config EnvVars, err error) {
	return EnvVars{
		AUTH0_DOMAIN:   os.Getenv("AUTH0_DOMAIN"),
		AUTH0_AUDIENCE: os.Getenv("AUTH0_AUDIENCE"),
		PORT:           os.Getenv("PORT"),
	}, err
}
