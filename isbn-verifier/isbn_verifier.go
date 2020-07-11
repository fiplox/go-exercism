package isbn

// IsValidISBN verifies book identification numbers.
func IsValidISBN(isbn string) bool {
	if len(isbn) != 13 {
		if len(isbn) != 10 {
			return false
		}
	}

	if isbn[len(isbn)-1] < 48 && isbn[len(isbn)-1] > 57 {
		if isbn[len(isbn)-1] != 'X' {
			return false
		}
	}

	res := 0
	j := 10
	for i, r := range isbn {
		if r == '-' {
			continue
		}

		if r == 'X' {
			if i != len(isbn)-1 {
				return false
			}
			res += 10
			break
		}
		res = (int(r-'0') * j) + res
		j--
	}

	return res%11 == 0
}
