package accumulate

// Accumulate runs provided funtion to each element of the given list
func Accumulate(s []string, fun func(string) string) []string {
	var res []string
	for _, ss := range s {
		res = append(res, fun(ss))
	}
	return res
}
