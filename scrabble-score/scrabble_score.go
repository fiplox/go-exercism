package scrabble

import "strings"

// Score counts the scrabble score for a given word
func Score(a string) int {
	a = strings.ToUpper(a)
	var res int

	for _, c := range a {
		switch c {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			res++
		case 'D', 'G':
			res += 2
		case 'B', 'C', 'M', 'P':
			res += 3
		case 'F', 'H', 'V', 'W', 'Y':
			res += 4
		case 'K':
			res += 5
		case 'J', 'X':
			res += 8
		case 'Q', 'Z':
			res += 10
		}
	}
	return res
}
