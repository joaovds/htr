package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/joaovds/htr/internal/config"
	"github.com/joaovds/htr/internal/ui"
)

type request struct {
	baseURL       string
	globalHeaders map[string]string
	reqConfig     config.Request
	args          map[string]string
	noStyle       bool
}

func New(baseURL string, reqConfig config.Request, globalHeaders map[string]string, args []string, noStyle bool) *request {
	return &request{baseURL, globalHeaders, reqConfig, parseArgs(args), noStyle}
}

func parseArgs(args []string) map[string]string {
	values := make(map[string]string)
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			continue
		}
		parts := strings.SplitN(arg, "=", 2)
		if len(parts) == 2 {
			values[parts[0]] = parts[1]
		}
	}
	return values
}

func (r *request) Run() error {
	err := r.applyReplacements()
	if err != nil {
		return err
	}

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

	for key, value := range r.globalHeaders {
		req.Header.Set(key, value)
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

func (r *request) applyReplacements() error {
	placeholders := r.getPlaceholders()

	for _, placeholder := range placeholders {
		if _, exists := r.args[placeholder]; !exists {
			return fmt.Errorf("missing required parameter: %s", placeholder)
		}
	}

	for key, value := range r.globalHeaders {
		r.globalHeaders[key] = r.replacePlaceholders(value)
	}
	for key, value := range r.reqConfig.Headers {
		r.reqConfig.Headers[key] = r.replacePlaceholders(value)
	}

	return nil
}

func (r *request) replacePlaceholders(input string) string {
	for key, value := range r.args {
		placeholder := fmt.Sprintf("#{{%s}}", key)
		input = strings.ReplaceAll(input, placeholder, value)
	}
	return input
}
