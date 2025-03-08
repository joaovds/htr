package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/joaovds/htr/internal/config"
)

func MakeRequest(baseURL string, reqConfig config.Request) error {
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
