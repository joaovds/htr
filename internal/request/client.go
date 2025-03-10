package request

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/joaovds/htr/internal/config"
	"github.com/joaovds/htr/internal/ui"
)

type request struct {
	baseURL   string
	reqConfig config.Request
	noStyle   bool
}

func New(baseURL string, reqConfig config.Request, noStyle bool) *request {
	return &request{baseURL, reqConfig, noStyle}
}

func (r *request) Run() error {
	client := &http.Client{}

	var bodyReader io.Reader
	if r.reqConfig.Body != nil {
		jsonBody, err := json.Marshal(r.reqConfig.Body)
		if err != nil {
			return err
		}
		bodyReader = strings.NewReader(string(jsonBody))
	}

	var url string
	if r.reqConfig.Url != "" {
		url = r.reqConfig.Url
	} else {
		if r.baseURL != "" && r.reqConfig.Endpoint != "" {
			url = r.baseURL + r.reqConfig.Endpoint
		} else {
			return errors.New("url or baseURL with required endpoint")
		}
	}

	req, err := http.NewRequest(r.reqConfig.Method, url, bodyReader)
	if err != nil {
		return err
	}

	for key, value := range r.reqConfig.Headers {
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
	var responseJSONStr string
	if json.Unmarshal(body, &prettyJSON) == nil {
		pretty, _ := json.MarshalIndent(prettyJSON, "", "  ")
		responseJSONStr = string(pretty)
	} else {
		responseJSONStr = string(body)
	}

	responseUI := ui.NewResponse(url, resp.StatusCode, responseJSONStr, r.noStyle)
	responseUI.Render()

	return nil
}
