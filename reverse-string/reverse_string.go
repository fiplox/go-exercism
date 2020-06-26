// Package reverse contains exercise reverse
package reverse

// Reverse reverses given string
func Reverse(s string) string {
	size := len(s)
	c := make([]rune, size)
	for _, rune := range s {
		size--
		c[size] = rune
	}
	return string(c[size:])
}
