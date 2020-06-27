package diffsquares

// SquareOfSum returns the square of the sum of first n numbers
/* 1 + 2 + 3 + ... + n = n(n + 1) / 2
   (n(n + 1) / 2)^2 = n^2(n+1)^2 / 4
*/
func SquareOfSum(n int) int {
	return n * n * (n*n + 2*n + 1) / 4
}

// SumOfSquares returns the sum of the squares of first n numbers
// 1^2 + 2^2 + 3^2 + ... + n^2 = n(n + 1)(2n + 1) / 6
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference returns the difference between
// square of the sum and the sum of the squares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
