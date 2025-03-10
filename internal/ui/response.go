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
		noStyle  bool
		sb       strings.Builder
	}
)

func NewResponse(endpoint string, code int, body string, noStyle bool) *Response {
	var sb strings.Builder
	return &Response{
		Endpoint: endpoint,
		HttpCode: code,
		Body:     body,
		noStyle:  noStyle,
		sb:       sb,
	}
}

func (r *Response) Render() {
	if r.noStyle {
		r.renderWithoutStyle()
		return
	}

	lines := SplitLines(r.Body)
	maxWidth := FindMaxWidth(lines)

	r.renderSeparator(maxWidth, DoubleTopLeftCorner, DoubleTopRightCorner)
	r.renderHttpCode(maxWidth)
	r.renderSeparator(maxWidth, DoubleMiddleLeftCorner, DoubleMiddleRightCorner)
	r.renderBody(lines, maxWidth)
	r.renderSeparator(maxWidth, DoubleBottomLeftCorner, DoubleBottomRightCorner)

	fmt.Print(r.sb.String())
}

func (r *Response) renderHttpCode(maxWidth int) {
	fmt.Fprintf(&r.sb, "%s Http Code: %-*d %s\n", DoubleVerticalLine, maxWidth-len("Http Code: "), r.HttpCode, DoubleVerticalLine)
}

func (r *Response) renderBody(lines []string, maxWidth int) {
	if len(lines) == 0 {
		fmt.Fprintf(&r.sb, "%s %-*s %s\n", DoubleVerticalLine, maxWidth, "No body data", DoubleVerticalLine)
		return
	}

	for _, line := range lines {
		fmt.Fprintf(&r.sb, "%s %-*s %s\n", DoubleVerticalLine, maxWidth, line, DoubleVerticalLine)
	}
}

func (r *Response) renderSeparator(width int, left, right string) {
	r.sb.WriteString(left + Repeat(DoubleHorizontalLine, width+2) + right + "\n")
}

func (r *Response) renderWithoutStyle() {
	fmt.Println("Http Code: ", r.HttpCode)
	fmt.Println("Body:")
	fmt.Println(r.Body)
}
