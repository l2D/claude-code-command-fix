package formatter

import (
	"regexp"
	"strings"
)

var (
	lineContinuationRe = regexp.MustCompile(`\\\s*\n\s*`)
	whitespaceRe       = regexp.MustCompile(`\s+`)
	andRe              = regexp.MustCompile(`\s*&&\s*`)
	pipeRe             = regexp.MustCompile(`\s*\|\s*`)
	semicolonRe        = regexp.MustCompile(`\s*;\s*`)
)

// FormatCommand cleans a terminal command by collapsing whitespace
// and normalizing operator spacing for &&, |, and ;.
// If the input contains backslash line continuations, they are
// preserved with consistent indentation.
func FormatCommand(commandText string, singleLine bool) string {
	text := strings.TrimSpace(commandText)
	if text == "" {
		return ""
	}

	if !singleLine && lineContinuationRe.MatchString(text) {
		return formatMultiLine(text)
	}

	text = lineContinuationRe.ReplaceAllString(text, " ")
	return formatSingleLine(text)
}

func formatSingleLine(text string) string {
	text = whitespaceRe.ReplaceAllString(text, " ")
	text = andRe.ReplaceAllString(text, " && ")
	text = pipeRe.ReplaceAllString(text, " | ")
	text = semicolonRe.ReplaceAllString(text, "; ")
	return text
}

func formatMultiLine(text string) string {
	parts := lineContinuationRe.Split(text, -1)
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}
	return strings.Join(parts, " \\\n  ")
}
