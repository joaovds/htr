package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		BaseURL  string             `yaml:"baseURL" json:"baseURL"`
		Requests map[string]Request `yaml:"requests" json:"requests"`
	}

	Request struct {
		Url      string            `yaml:"url" json:"url"`
		Endpoint string            `yaml:"endpoint" json:"endpoint"`
		Method   string            `yaml:"method" json:"method"`
		Body     any               `yaml:"body" json:"body"`
		Headers  map[string]string `yaml:"headers" json:"headers"`
	}
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
