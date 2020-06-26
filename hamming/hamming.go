package hamming

import "errors"

// Distance finds the distance between two DNA strands
func Distance(a, b string) (int, error) {
	l := 0

	if len(a) != len(b) {
		return l, errors.New("a != b")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			l++
		}
	}
	return l, nil
}
