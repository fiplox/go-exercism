// Package acronym contains exercise acronym.
package acronym

import "strings"

// Abbreviate converts phare into acronym.
func Abbreviate(s string) string {

	var res string
	ss := strings.Fields(strings.ReplaceAll(strings.ReplaceAll(s, "-", " "), "_", " "))
	for _, w := range ss {
		res += string(w[0])
	}
	return strings.ToUpper(res)
}
