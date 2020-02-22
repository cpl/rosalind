package rosalind

func ProteinToRNACount(protein string) int {
	count := 1
	for _, symbol := range protein {
		count *= len(CodonRNATable[symbol])
		count %= 1000000
	}
	return count
}

var CodonRNATable = map[rune][]string{
	0:   {"UAA", "UAG", "UGA"},
	'G': {"GGU", "GGC", "GGA", "GGG"},
	'E': {"GAA", "GAG"},
	'D': {"GAU", "GAC"},
	'A': {"GCU", "GCC", "GCA", "GCG"},
	'V': {"GUU", "GUC", "GUA", "GUG"},
	'R': {"AGA", "AGG", "CGU", "CGC", "CGA", "CGG"},
	'S': {"AGU", "AGC", "UCU", "UCC", "UCA", "UCG"},
	'K': {"AAA", "AAG"},
	'N': {"AAU", "AAC"},
	'T': {"ACG", "ACA", "ACC", "ACU"},
	'M': {"AUG"},
	'I': {"AUU", "AUC", "AUA"},
	'Q': {"CAA", "CAG"},
	'H': {"CAU", "CAC"},
	'P': {"CCG", "CCA", "CCC", "CCU"},
	'L': {"CUU", "CUC", "CUA", "CUG", "UUA", "UUG"},
	'W': {"UGG"},
	'C': {"UGU", "UGC"},
	'Y': {"UAC", "UAU"},
	'F': {"UUU", "UUC"},
}
