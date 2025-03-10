package ui

import (
	"fmt"
	"strings"
)

type (
	Response struct {
		Endpoint string
		HttpCode int
		Body     string
	}
)

func NewResponse(endpoint string, code int, body string) *Response {
	return &Response{
		Endpoint: endpoint,
		HttpCode: code,
		Body:     body,
	}
}

func (r *Response) Render() {
	lines := SplitLines(r.Body)
	maxWidth := FindMaxWidth(lines)
	var sb strings.Builder

	renderSeparator(&sb, maxWidth, DoubleTopLeftCorner, DoubleTopRightCorner)
	renderHttpCode(&sb, r.HttpCode, maxWidth)
	renderSeparator(&sb, maxWidth, DoubleMiddleLeftCorner, DoubleMiddleRightCorner)
	renderBody(&sb, lines, maxWidth)
	renderSeparator(&sb, maxWidth, DoubleBottomLeftCorner, DoubleBottomRightCorner)

	fmt.Print(sb.String())
}

func renderHttpCode(sb *strings.Builder, code, maxWidth int) {
	fmt.Fprintf(sb, "%s Http Code: %-*d %s\n", DoubleVerticalLine, maxWidth-len("Http Code: "), code, DoubleVerticalLine)
}

func renderBody(sb *strings.Builder, lines []string, maxWidth int) {
	if len(lines) == 0 {
		fmt.Fprintf(sb, "%s %-*s %s\n", DoubleVerticalLine, maxWidth, "No body data", DoubleVerticalLine)
		return
	}

	for _, line := range lines {
		fmt.Fprintf(sb, "%s %-*s %s\n", DoubleVerticalLine, maxWidth, line, DoubleVerticalLine)
	}
}

func renderSeparator(sb *strings.Builder, width int, left, right string) {
	sb.WriteString(left + Repeat(DoubleHorizontalLine, width+2) + right + "\n")
}
