package triangle

import "math"

// Kind is a kind of a triangle
type Kind string

const (
	// NaT - not a triangle.
	NaT = Kind("NaT")
	// Equ - equilateral.
	Equ = Kind("Equ")
	// Iso - isosceles.
	Iso = Kind("Iso")
	// Sca - scalene.
	Sca = Kind("Sca")
)

// valid returns true if a, b and c makes a triangle.
func valid(a, b, c float64) bool {
	if a+b < c || b+c < a || c+a < b {
		return false
	}
	if a <= 0 || math.IsNaN(a) || math.IsInf(a, 0) {
		return false
	}
	if b <= 0 || math.IsNaN(b) || math.IsInf(b, 0) {
		return false
	}
	if c <= 0 || math.IsNaN(c) || math.IsInf(c, 0) {
		return false
	}
	return true
}

// KindFromSides return the type of a triangle by given 3 sides length.
func KindFromSides(a, b, c float64) Kind {
	if !valid(a, b, c) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || b == c || a == c {
		return Iso
	}
	return Sca
}
