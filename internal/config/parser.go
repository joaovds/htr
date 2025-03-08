package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig(filepath string) (*Config, error) {
	if filepath == "" {
		return nil, errors.New("you need the path to the file")
	}

	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var config Config
	if err := yaml.Unmarshal(file, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
