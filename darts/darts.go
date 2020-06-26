// Package darts contains darts exercise
package darts

import "math"

// Score returns the darts score by given (x,y) coordinates
func Score(x, y float64) int {
	r := math.Sqrt(x*x + y*y)
	switch {
	case r <= 1:
		return 10
	case r <= 5:
		return 5
	case r <= 10:
		return 1
	}
	return 0
}
