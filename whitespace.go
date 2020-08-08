package hagl

func collapseWhitespace(text string) string {
	b := make([]byte, 0, len(text))

	lastWasSpace := false
	for i := 0; i < len(text); i++ {
		c := text[i]
		isWhitespace := c == '\n' || c == ' ' || c == '\t' || c == '\r'

		// Don't add consecutive spaces
		if isWhitespace && lastWasSpace {
			continue
		}

		// Convert all whitespace to a single space
		if isWhitespace {
			c = ' '
		}

		b = append(b, c)
		lastWasSpace = isWhitespace
	}

	return string(b)
}
