package rosalind

func RNAtoProtein(rna string) string {
	protein := make([]rune, len(rna)/3)

	for index := 0; index < len(rna)/3; index++ {
		symbol := RNACodonTable[rna[index*3:(index+1)*3]]
		if symbol == 0 {
			return string(protein[:index])
		}
		protein[index] = symbol
	}
	return string(protein)
}

var RNACodonTable = map[string]rune{
	"UUU": 'F',
	"UUC": 'F',
	"UUA": 'L',
	"UUG": 'L',
	"UCU": 'S',
	"UCC": 'S',
	"UCA": 'S',
	"UCG": 'S',
	"UAU": 'Y',
	"UAC": 'Y',
	"UAA": 0,
	"UAG": 0,
	"UGU": 'C',
	"UGC": 'C',
	"UGA": 0,
	"UGG": 'W',

	"CUU": 'L',
	"CUC": 'L',
	"CUA": 'L',
	"CUG": 'L',
	"CCU": 'P',
	"CCC": 'P',
	"CCA": 'P',
	"CCG": 'P',
	"CAU": 'H',
	"CAC": 'H',
	"CAA": 'Q',
	"CAG": 'Q',
	"CGU": 'R',
	"CGC": 'R',
	"CGA": 'R',
	"CGG": 'R',

	"AUU": 'I',
	"AUC": 'I',
	"AUA": 'I',
	"AUG": 'M',
	"ACU": 'T',
	"ACC": 'T',
	"ACA": 'T',
	"ACG": 'T',
	"AAU": 'N',
	"AAC": 'N',
	"AAA": 'K',
	"AAG": 'K',
	"AGU": 'S',
	"AGC": 'S',
	"AGA": 'R',
	"AGG": 'R',

	"GUU": 'V',
	"GUC": 'V',
	"GUA": 'V',
	"GUG": 'V',
	"GCU": 'A',
	"GCC": 'A',
	"GCA": 'A',
	"GCG": 'A',
	"GAU": 'D',
	"GAC": 'D',
	"GAA": 'E',
	"GAG": 'E',
	"GGU": 'G',
	"GGC": 'G',
	"GGA": 'G',
	"GGG": 'G',
}
