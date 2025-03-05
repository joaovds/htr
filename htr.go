package htr

import (
	"os"

	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		Requests map[string]Request `yaml:"requests" json:"requests"`
	}

	Request struct {
		Url    string `yaml:"url" json:"url"`
		Method string `yaml:"method" json:"method"`
	}
)

func LoadConfig(filename string) (*Config, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var config Config
	if err := yaml.Unmarshal(file, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
