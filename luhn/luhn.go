package luhn

import "strings"

// Valid determines whether or not s is valid per the Luhn formula.
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")

	if len(input) <= 1 {
		return false
	}

	var sum int
	var digit int

	double := len(input)%2 == 0

	for _, r := range input {
		digit = int(r - '0')
		if digit < 0 || digit > 9 {
			return false
		}
		if double {
			digit *= 2
			if digit > 9 {
				sum -= 9
			}
		}
		sum += digit
		double = !double
	}

	return sum%10 == 0
}
