package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

type HTTP struct {
	Port string `env:"HTTP_PORT" env-default:"8080"`
}

type DB struct {
	Name     string `env:"DB_NAME" env-default:"postgres"`
	Host     string `env:"DB_HOST" env-default:"postgres"`
	Port     string `env:"DB_PORT" env-default:"5432"`
	User     string `env:"DB_USER" env-default:"postgres"`
	Password string `env:"DB_PASSWORD" env-default:"postgres"`
}

type Tokens struct {
	User  string `env:"USER_TOKEN" env-default:"user_token"`
	Admin string `env:"ADMIN_TOKEN" env-default:"admin_token"`
}

type Cache struct {
	Host string `env:"REDIS_HOST" env-default:"redis"`
	Port string `env:"REDIS_PORT" env-default:"6379"`
}

type Config struct {
	HTTP
	DB
	Cache
	Tokens
}

func New() (*Config, error) {

	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, errors.Wrap(err, "reading config error")
	}

	return &cfg, nil
}
