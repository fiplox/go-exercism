package atbash

import "strings"

// NotValid returns true if rune contains ,-'. or space.
func NotValid(r rune) bool {
	switch r {
	case ' ', ',', '-', '\'', '.':
		return true
	}
	return false
}

// Atbash encodes string into atbash cipher.
func Atbash(s string) string {
	var cipher string

	if len(s) == 0 {
		return cipher
	}
	s = strings.ToLower(s)
	i := 0
	for _, r := range s {
		if NotValid(r) {
			continue
		}
		if i%5 == 0 && i != 0 {
			cipher += " "
		}
		if r >= 97 && r <= 122 {
			cipher += string('a' + 'z' - r)
		} else {
			cipher += string(r)
		}
		i++
	}

	return cipher
}
