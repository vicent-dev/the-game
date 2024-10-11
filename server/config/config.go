package config

import (
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
}

func LoadConfig() *Config {
	c := &Config{}

	cFile := getConfigFile()
	err := yaml.Unmarshal(cFile, c)

	if err != nil {
		log.Fatalln(err)
	}

	return c
}
