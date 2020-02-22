package rosalind

func ReverseComplement(dna string) string {
	complement := make([]rune, len(dna))
	for index, symbol := range dna {
		complement[len(dna)-index-1] = Complement(symbol)
	}
	return string(complement)
}

func Complement(symbol rune) rune {
	switch symbol {
	case 'A':
		return 'T'
	case 'C':
		return 'G'
	case 'G':
		return 'C'
	case 'T':
		return 'A'
	default:
		panic("bad symbol")
	}
}
