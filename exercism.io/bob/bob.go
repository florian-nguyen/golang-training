// Package bob generates the replies of the lackadaisical teenager Bob.
package bob

// Import packages.
import (
	"strings"
	"unicode"
)

// hasAny tests a string expression with a function as criterion.
func hasAny(s string, testFunc func(rune) bool) bool {
	for _, c := range s {
		if testFunc(c) {
			return true
		}
	}
	return false
}

// Hey generates Bob's answer to any string remark.
func Hey(remark string) string {

	// Start by removing spaces from the string.
	remark = strings.TrimSpace(remark)

	// List of Bob's possible reactions.
	switch {
	case remark == "":
		return "Fine. Be that way!" // Nothing.
	case !strings.HasSuffix(remark, "?") && hasAny(remark, unicode.IsUpper) && !hasAny(remark, unicode.IsLower):
		return "Whoa, chill out!" // Not a question, shouted.
	case strings.HasSuffix(remark, "?") && hasAny(remark, unicode.IsUpper) && !hasAny(remark, unicode.IsLower):
		return "Calm down, I know what I'm doing!" // Question and shouted.
	case strings.HasSuffix(remark, "?") && !(hasAny(remark, unicode.IsUpper) && !hasAny(remark, unicode.IsLower)):
		return "Sure." // Question, not shouted.
	default:
		return "Whatever." // Other cases.
	}
}
