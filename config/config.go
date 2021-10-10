package config

import (
	"os"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func Init() *Config {
	return &Config{
		DBUsername: os.Getenv("PG_USER"),
		DBPassword: os.Getenv("PG_PASS"),
		DBHost:     os.Getenv("PG_HOST"),
		DBPort:     os.Getenv("PG_PORT"),
		DBName:     os.Getenv("PG_NAME"),
	}
}

func (c *Config) Env(name string) string {
	return os.Getenv(name)
}
