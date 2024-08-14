package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env-default:"prod"`
	Version     string `yaml:"version" env-default:"0.0.0"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	Server      `yaml:"server"`
}

type Server struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"30s"`
}

func New() *Config {
	// Get the path to the configuration file from the environment
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// Check if the configuration file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Configuration file %s does not exist", configPath)
	}

	// Read the configuration file
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Failed to read configuration: %v", err)
	}

	return &cfg
}
