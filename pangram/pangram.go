package pangram

import "strings"

// IsPangram determines if a sentence is a pangram.
func IsPangram(s string) bool {
	if s == "" {
		return false
	}

	s = strings.ToLower(s)
	m := map[rune]bool{}
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for _, letter := range s {
		m[letter] = true
	}

	for _, a := range alphabet {
		if !m[a] {
			return false
		}
	}
	return true
}
