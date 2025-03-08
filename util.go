package hagl

import (
	"regexp"
)

func sanitizeAttrName(attr string) string {
	// Allow only a-z, A-Z, 0-9, and - (dash)
	re := regexp.MustCompile(`[^a-zA-Z0-9-]`)
	return re.ReplaceAllString(attr, "")
}
