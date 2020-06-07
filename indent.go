package go_temper

import (
	"strings"
)

func indent(n int, text string, tab string) string {
	prefix := ""
	for i := 0; i < n; i++ {
		prefix += tab
	}

	lines := strings.Split(text, "\n")

	for i, line := range lines {
		lines[i] = prefix + line
	}

	return strings.Join(lines, "\n")
}
