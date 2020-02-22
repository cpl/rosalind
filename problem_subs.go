package rosalind

import "strings"

func Motif(dna, substring string) []int {
	indices := make([]int, 0)

	for index := range dna {
		if strings.HasPrefix(dna[index:], substring) {
			indices = append(indices, index+1)
		}
	}

	return indices
}
