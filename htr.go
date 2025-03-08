package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

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

func main() {
	filename := os.Args[1]
	config, err := LoadConfig(filename)
	if err != nil {
		fmt.Println("Err file load: ", err)
		os.Exit(1)
	}

	if len(os.Args) < 3 {
		fmt.Println("Requests Available:")
		for name := range config.Requests {
			fmt.Println("-", name)
		}
		os.Exit(0)
	}

	reqName := os.Args[2]
	reqConfig, exists := config.Requests[reqName]
	if !exists {
		fmt.Println("Request not found:", reqName)
		os.Exit(1)
	}

	if err := makeRequest(config.BaseURL, reqConfig); err != nil {
		fmt.Println("Request fail:", err)
		os.Exit(1)
	}
}

func makeRequest(baseURL string, reqConfig Request) error {
	client := &http.Client{}

	var bodyReader io.Reader
	if reqConfig.Body != nil {
		jsonBody, err := json.Marshal(reqConfig.Body)
		if err != nil {
			return err
		}
		bodyReader = strings.NewReader(string(jsonBody))
	}

	var url string
	if reqConfig.Url != "" {
		url = reqConfig.Url
	} else {
		if baseURL != "" && reqConfig.Endpoint != "" {
			url = baseURL + reqConfig.Endpoint
		} else {
			return errors.New("url or baseURL with required endpoint")
		}
	}

	req, err := http.NewRequest(reqConfig.Method, url, bodyReader)
	if err != nil {
		return err
	}

	for key, value := range reqConfig.Headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var prettyJSON map[string]any
	fmt.Println("HttpCode:", resp.StatusCode)
	if json.Unmarshal(body, &prettyJSON) == nil {
		pretty, _ := json.MarshalIndent(prettyJSON, "", "  ")
		fmt.Println(string(pretty))
	} else {
		fmt.Println(string(body))
	}

	return nil
}
