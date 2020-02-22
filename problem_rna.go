package rosalind

import "strings"

func DNAtoRNA(dna string) string {
	return strings.ReplaceAll(dna, "T", "U")
}
