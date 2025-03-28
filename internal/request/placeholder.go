package request

import "regexp"

func (r *request) getPlaceholders() []string {
	var placeholders []string
	placeholderMap := make(map[string]any)

	extractPlaceholders := func(text string) {
		regex := regexp.MustCompile("#{{(.*?)}}")
		matches := regex.FindAllStringSubmatch(text, -1)
		for _, match := range matches {
			if len(match) > 1 {
				key := match[1]
				if _, exists := placeholderMap[key]; !exists {
					placeholderMap[key] = struct{}{}
					placeholders = append(placeholders, key)
				}
			}
		}
	}

	for _, value := range r.globalHeaders {
		extractPlaceholders(value)
	}

	for _, value := range r.reqConfig.Headers {
		extractPlaceholders(value)
	}

	return placeholders
}
