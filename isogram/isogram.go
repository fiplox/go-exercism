// Package isogram contains exercise isogram
package isogram

import "strings"

// IsIsogram verifies if given word/phare is isogram
func IsIsogram(s string) bool {
	s = strings.ToUpper(s)
	var n int
	for _, ss := range s {
		if ss == ' ' || ss == '-' {
			continue
		}
		n = strings.Count(s, string(ss))
		if n > 1 {
			return false
		}
	}
	return true
}
