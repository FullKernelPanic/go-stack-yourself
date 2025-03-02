package services

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App struct {
		Name string `yaml:"name"`
		Web  struct {
			Port     int    `yaml:"port"`
			BasePath string `yaml:"basepath"`
		}
	} `yaml:"app"`
	Infrastructure struct {
		Monitoring struct {
			Enabled           bool   `yaml:"enabled"`
			OtelCollectorHost string `yaml:"otel-collector-host"`
		} `yaml:"monitoring"`
	}
}

var (
	configPath     string
	configInstance *Config
)

func GetConfig() *Config {
	if configInstance != nil {
		return configInstance
	}

	configInstance = loadConfig()
	return configInstance
}

func SetConfigPath(path string) {
	configPath = path
}

func loadConfig() *Config {
	yamlFile, err := os.ReadFile(configPath)

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	c := &Config{}
	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
