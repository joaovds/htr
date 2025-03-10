package ui

import (
	"bufio"
	"fmt"
	"strings"
)

func MoveCursor(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col)
}

func SplitLines(s string) []string {
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func FindMaxWidth(lines []string) int {
	maxWidth := 20
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}
	return maxWidth
}

func Repeat(char string, count int) string {
	var sb strings.Builder
	sb.Grow(count)
	for range count {
		sb.WriteString(char)
	}
	return sb.String()
}
