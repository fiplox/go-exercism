package twofer // import "twofer"

import "strings"

//ShareWith returns the string with the given name or "you"if name is empty.
func ShareWith(name string) string {
	share := "One for you, one for me."

	if name == "" {
		return share
	}

	return strings.Replace(share, "you", name, 1)
}
