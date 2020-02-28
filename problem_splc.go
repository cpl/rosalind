package rosalind

import "strings"

func Splice(dna string, substrings ...string) string {
	for _, subs := range substrings {
		dna = strings.ReplaceAll(dna, subs, "")
	}
	return dna
}

func DNAtoProtein(dna string) string {
	return RNAtoProtein(DNAtoRNA(dna))
}
