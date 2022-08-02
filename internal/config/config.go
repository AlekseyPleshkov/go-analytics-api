package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	configFile = "../config.yml"
)

func NewConfig() (*Config, error) {
	data, err := ioutil.ReadFile(configFile)

	if err != nil {
		return nil, err
	}

	config := Config{}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
