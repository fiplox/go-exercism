// Package collatzconjecture contains exercise Collatz Conjecture
package collatzconjecture

import "errors"

// CollatzConjecture returns number of steps required to reach 1
func CollatzConjecture(n int) (int, error) {
	var steps int
	if n <= 0 {
		return n, errors.New("Negative or zero")
	}
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
		steps++
	}
	return steps, nil
}
