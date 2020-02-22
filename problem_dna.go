package rosalind

// http://rosalind.info/problems/dna/

func CountACGT(dna string) map[rune]int {
	ret := map[rune]int{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	for _, symbol := range dna {
		ret[symbol] += 1
	}

	return ret
}
