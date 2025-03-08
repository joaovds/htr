package config

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
