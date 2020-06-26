package romannumerals

import "errors"

// ToRomanNumeral converts number into roman numerals
func ToRomanNumeral(n int) (string, error) {
	if n <= 0 || n > 3000 {
		return "", errors.New("out of range")
	}

	var res string
	var match = []struct {
		a int
		r string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _, val := range match {
		for n >= val.a {
			res += val.r
			n -= val.a
		}
	}

	return res, nil
}
