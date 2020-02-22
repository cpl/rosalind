package rosalind

func HammingDistance(dna0, dna1 string) int {
	if len(dna0) != len(dna1) {
		return -1
	}

	diff := 0
	for index := range dna0 {
		if dna0[index] != dna1[index] {
			diff++
		}
	}
	return diff
}
