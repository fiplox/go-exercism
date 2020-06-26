package etl

import "strings"

// Transform transforms old scrabble to new scrabble format
func Transform(old map[int][]string) map[string]int {
	newScrabble := map[string]int{}

	for i, sa := range old {
		for _, s := range sa {
			newScrabble[strings.ToLower(s)] = i
		}
	}
	return newScrabble
}
