package formatter

import (
	"regexp"
	"strings"
)

var (
	whitespaceRe = regexp.MustCompile(`\s+`)
	andRe        = regexp.MustCompile(`\s*&&\s*`)
	pipeRe       = regexp.MustCompile(`\s*\|\s*`)
	semicolonRe  = regexp.MustCompile(`\s*;\s*`)
)

// FormatCommand cleans a terminal command by collapsing whitespace
// and normalizing operator spacing for &&, |, and ;.
func FormatCommand(commandText string) string {
	text := strings.TrimSpace(commandText)
	if text == "" {
		return ""
	}

	text = whitespaceRe.ReplaceAllString(text, " ")
	text = andRe.ReplaceAllString(text, " && ")
	text = pipeRe.ReplaceAllString(text, " | ")
	text = semicolonRe.ReplaceAllString(text, "; ")

	return text
}
