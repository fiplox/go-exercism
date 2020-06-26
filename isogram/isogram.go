// Package isogram contains exercise isogram
package isogram

import "strings"

// Once returns true if a given rune is present once in a string
func Once(s string, c rune) bool {
	n := 0
	for _, ss := range s {
		if ss == ' ' || ss == '-' {
			continue
		}
		if c == ss {
			n++
		}

		if n > 1 {
			return false
		}
	}

	return true
}

// IsIsogram verifies if given word/phare is isogram
func IsIsogram(s string) bool {
	s = strings.ToUpper(s)
	for _, ss := range s {
		if !Once(s, ss) {
			return false
		}
	}
	return true
}
