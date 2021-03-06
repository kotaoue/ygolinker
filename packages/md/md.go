package md

import (
	"fmt"
	"regexp"
	"strings"
)

// listMarkers is list marker by markdown.
func listMarkers() []string {
	// Sorting is important. '- [x]' and '- [ ]' must be ahead of '-'.
	s := []string{"- [x]", "- [ ]", "-", "*"}
	return s
}

// IsList is determine if the string is a list.
func IsList(s string) bool {
	s = strings.TrimSpace(s)

	for _, v := range listMarkers() {
		if strings.HasPrefix(s, fmt.Sprintf("%s ", v)) {
			return true
		}
	}

	return false
}

// IsLink is determine if the string is a link.
func IsLink(s string) bool {
	s = strings.TrimSpace(s)
	r := regexp.MustCompile(`\[.*\]\(http.*\)`)
	return r.MatchString(s)
}

// ListText is extract a string from the list.
func ListText(s string) string {
	s = strings.TrimSpace(s)

	for _, v := range listMarkers() {
		s = strings.Replace(s, fmt.Sprintf("%s ", v), "", 1)
	}

	return s
}
