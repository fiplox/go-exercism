// Package proverb ...
package proverb

// Proverb generates a proverb from a string list.
func Proverb(rhyme []string) []string {
	var s int = len(rhyme)
	res := []string{}

	for i := 0; i < s; i++ {
		if i < s-1 {
			res = append(res, "For want of a "+rhyme[i]+" the "+rhyme[i+1]+" was lost.")
		} else {
			res = append(res, "And all for the want of a "+rhyme[0]+".")
		}
	}

	return res
}
