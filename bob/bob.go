// Package bob contains string exercise
package bob

import "strings"

// hasLetter checks if string has at least one letter
func hasLetter(str string) bool {
	for _, c := range str {
		if (c > 'a' && c < 'z') || (c > 'A' && c < 'Z') {
			return true
		}
	}
	return false
}

// Hey returns Bob's response
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	if remark[len(remark)-1:] == "?" {
		if remark == strings.ToUpper(remark) {
			if !hasLetter(remark) {
				return "Sure."
			}
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}

	if remark == strings.ToUpper(remark) {
		if !hasLetter(remark) {
			return "Whatever."
		}
		return "Whoa, chill out!"
	}

	return "Whatever."
}
