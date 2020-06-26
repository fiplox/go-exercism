// Package protein contains exercise from exercism
package protein

import "errors"

var (
	// ErrStop tells if STOP codon is found
	ErrStop = errors.New("stop codon found")
	// ErrInvalidBase reports if codon does not exist
	ErrInvalidBase = errors.New("invalid codon")
)

// FromCodon translates codon to a protein name
func FromCodon(s string) (string, error) {
	switch s {
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	}
	return "", ErrInvalidBase
}

// FromRNA translates RNA sequence into proteins
func FromRNA(s string) ([]string, error) {
	var res []string
	var tmp string
	var err error
	for i := 0; i < len(s)-2; i += 3 {
		tmp, err = FromCodon(s[i:(i + 3)])
		if err == nil {
			res = append(res, tmp)
		} else if err == ErrStop {
			return res, nil
		} else {
			return res, err
		}
	}
	return res, err
}
