package grains

import "errors"

// Square calculates the number of grains of wheat
// on a chessboard given that the number on each square doubles.
func Square(square int) (uint64, error) {
	if square < 1 || square > 64 {
		return 0, errors.New("out of range")
	}

	return 1 << (square - 1), nil
}

// Total calculates the total number of grains on the chessboard.
func Total() uint64 {
	return 1<<64 - 1
}
