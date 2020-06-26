package strain

// Ints is a collection of integers
type Ints []int

// Lists is a collection of integer arrays
type Lists [][]int

// Strings is a collection of strings
type Strings []string

// Keep returns new collection of Ints if fun is true
func (x Ints) Keep(fun func(int) bool) Ints {
	var res Ints
	for _, n := range x {
		if fun(n) {
			res = append(res, n)
		}
	}
	return res
}

// Discard returns new collection of Ints if fun is not true
func (x Ints) Discard(fun func(int) bool) Ints {
	return x.Keep(func(y int) bool { return !fun(y) })
}

// Keep returns new collection if Lists if fun is true
func (l Lists) Keep(fun func([]int) bool) Lists {
	var res Lists
	for _, list := range l {
		if fun(list) {
			res = append(res, list)
		}
	}
	return res
}

// Keep returns new collection of Strings if fun is true
func (s Strings) Keep(fun func(string) bool) Strings {
	var res Strings
	for _, ss := range s {
		if fun(ss) {
			res = append(res, ss)
		}
	}
	return res
}
