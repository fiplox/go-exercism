package lsproduct

import (
	"errors"
	"strconv"
)

// LargestSeriesProduct calculates the largest product for a 
// contiguous substring of digits of length n.
func LargestSeriesProduct(digits string, span int) (int, error) {
	if len(digits) < span {
		return -1, errors.New("span must be smaller than string length")
	}
	if span < 0 {
		return -1, errors.New("span must be grater that zero")
	}
	var product int
	for i := 0; i <= len(digits) - span; i++ {
		test := 1
		for j:= 0; j < span; j++ {
			digit, err := strconv.Atoi(string(digits[i+j]))
			if err != nil {
				return -1, errors.New("digits input must only contain digits")
			}
			test *= digit
		}
		if test > product {
			product = test
		}
	}
	return product, nil
}
