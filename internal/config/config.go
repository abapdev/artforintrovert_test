package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Listen struct {
		Paths struct {
			Base string `yaml:"base" env:"APP_BASE_PATH" env-default:"/testapp" env-description:"Base path for all endpoints"`
		} `yaml:"paths"`
		Ports struct {
			Main string `yaml:"main" env:"APP_HTTP_PORT" env-default:"3000" env-description:"HTTP port for http endpoint"`
		} `yaml:"ports"`
	} `yaml:"listen"`
	Storage struct {
		Mongo struct {
			Connect string `yaml:"connect" env:"MONGO_PATH" env-default:"/" env-description:"Path to MongoDB"`
		} `yaml:"mongo"`
	} `yaml:"storage"`
}

func ReadConfig(configPath string) (*Config, error) {
	var config Config
	var err error

	if configExists(configPath) {
		err = cleanenv.ReadEnv(&config)
	} else {
		err = cleanenv.ReadConfig(configPath, &config)
	}

	if err != nil {
		return nil, err
	}

	return &config, nil
}

func configExists(configPath string) bool {
	_, err := os.Stat(configPath)
	return os.IsNotExist(err)
}
