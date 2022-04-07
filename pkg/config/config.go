package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// Config represents mix of settings for the app.
type Config struct {
	Server Server  `yaml:"server"`
	Mongo  MongoDB `yaml:"mongo"`
}

// New creates reads application configuration from the file.
func New(path string) (*Config, error) {
	var config Config
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(file, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
