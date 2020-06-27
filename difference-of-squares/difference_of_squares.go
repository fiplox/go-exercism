package diffsquares

// SquareOfSum returns the square of the sum of first n numbers
func SquareOfSum(n int) int {
	var res int

	for i := 0; i <= n; i++ {
		res += i
	}

	return res * res
}

// SumOfSquares returns the sum of the squares of first n numbers
func SumOfSquares(n int) int {
	var res int

	for i := 0; i <= n; i++ {
		res += i * i
	}

	return res
}

// Difference returns the difference between
// square of the sum and the sum of the squares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
