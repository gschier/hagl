package hagl

import (
	"strings"
)

func WrapText(text string, lineLength int) string {
	words := strings.Split(text, " ")
	var wrappedLines []string
	currentLine := ""

	for i, word := range words {
		if len(currentLine)+len(word)+1 > lineLength {
			wrappedLines = append(wrappedLines, currentLine)
			currentLine = ""
		}
		currentLine += word
		if i != len(words)-1 {
			currentLine += " "
		}
	}

	if currentLine != "" {
		wrappedLines = append(wrappedLines, currentLine)
	}

	return strings.Join(wrappedLines, "\n")
}
