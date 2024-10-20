package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port int `envconfig:"PORT" default:"8888"`
}

func LoadConfig[T any](cfg *T) error {
	err := envconfig.Process("", cfg)
	if err != nil {
		return err
	}

	return nil
}
