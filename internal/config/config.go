package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string     `yaml:"env" toml:"env" env-default:"local"`
	DBConfig   DBConfig   `yaml:"database" toml:"database"`
	HTTPServer HTTPServer `yaml:"http_server" toml:"http_server"`
}

type DBConfig struct {
	Name     string `yaml:"name" toml:"name" env-default:"postgres"`
	Host     string `yaml:"host" toml:"host" env-default:"postgres"`
	Port     string `yaml:"port" toml:"port" env-default:"5432"`
	User     string `yaml:"user" toml:"user" env-default:"postgres"`
	Password string `yaml:"password" toml:"password" env-default:"postgres"`
	DBName   string `yaml:"db_name" toml:"db_name" env-default:"postgres"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" toml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" toml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" toml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %v", err)
	}

	return &cfg
}
