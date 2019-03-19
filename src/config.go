package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)

type (
	ConfigInterface interface {
		LoadConfig(file string)
		ParseConfig(data []byte) error
		IsDebug() bool
	}

	HttpSection struct {
		Host string `yml: "host"`
		Port int    `yml: "port"`
	}

	Config struct {
		ConfigInterface
		Debug bool	`yml: "debug"`
		Http HttpSection `yml: "http"`
	}
)

func (config *Config) LoadConfig(file string) {
	filename, _ := filepath.Abs(file)
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalf("Load config: %s", err)
	}

	err = config.ParseConfig(data)

	if err != nil {
		log.Fatalf("Parse config: %s", err)
	}
}

func (config *Config) ParseConfig(data []byte) error {
	return yaml.Unmarshal(data, &config)
}

func (config *Config) IsDebug() bool {
	return config.Debug
}
