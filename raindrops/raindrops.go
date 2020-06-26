package raindrops

import "strconv"

//Convert converts a number into a specific string
//if it corresponds to certain potential factors.
func Convert(n int) string {
	if n%3 != 0 && n%5 != 0 && n%7 != 0 {
		return strconv.Itoa(n)
	}

	var res string
	if n%3 == 0 {
		res += "Pling"
	}
	if n%5 == 0 {
		res += "Plang"
	}
	if n%7 == 0 {
		res += "Plong"
	}

	return res
}
