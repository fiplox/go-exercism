package strand

import "strings"

//replacer replaces each neucleotide with its complement
var replacer = strings.NewReplacer("C", "G", "G", "C", "A", "U", "T", "A")

// ToRNA transforms DNA strand to its RNA complement
func ToRNA(dna string) string {
	if dna == "" {
		return dna
	}
	dna = replacer.Replace(dna)
	return dna
}
