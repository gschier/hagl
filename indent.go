package go_temper

import (
	"strings"
)

func indent(level int, text string, tab string) string {
	return indentN(level, text, tab, -1)
}

func indentN(level int, text string, tab string, n int) string {
	prefix := ""
	for i := 0; i < level; i++ {
		prefix += tab
	}

	lines := strings.Split(text, "\n")

	for i, line := range lines {
		if n >= 0 && i >= n {
			break
		}

		lines[i] = prefix + line
	}

	return strings.Join(lines, "\n")
}
