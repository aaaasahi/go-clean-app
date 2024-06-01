package config

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	Conf struct {
		DB DB
	}
	DB struct {
		Host     string `envconfig:"DB_HOST"`
		Port     string `envconfig:"DB_PORT"`
		User     string `envconfig:"DB_USER"`
		Password string `envconfig:"DB_PASSWORD"`
		Name     string `envconfig:"DB_NAME"`
	}
)

func (c *Conf) Init() error {
	err := envconfig.Process("", c)
	if err != nil {
		return err
	}
	return nil
}