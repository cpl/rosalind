package rosalind

import "fmt"

var mapAminoAcids = map[string]uint8{
	// DNA
	"TTT": 'F',
	"TTC": 'F',
	"TTA": 'L',
	"TTG": 'L',
	"TCT": 'S',
	"TCC": 'S',
	"TCA": 'S',
	"TCG": 'S',
	"TAT": 'Y',
	"TAC": 'Y',
	"TAA": '_',
	"TAG": '_',
	"TGT": 'C',
	"TGC": 'C',
	"TGA": '_',
	"TGG": 'W',
	"CTT": 'L',
	"CTC": 'L',
	"CTA": 'L',
	"CTG": 'L',
	"CCT": 'P',
	"CAT": 'H',
	"CGT": 'R',
	"ATT": 'I',
	"ATC": 'I',
	"ATA": 'I',
	"ATG": 'M',
	"ACT": 'T',
	"AAT": 'N',
	"AGT": 'S',
	"GTT": 'V',
	"GTC": 'V',
	"GTA": 'V',
	"GTG": 'V',
	"GCT": 'A',
	"GAT": 'D',
	"GGT": 'G',

	// RNA
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
	"UAA": '_',
	"UAG": '_',
	"UGU": 'C',
	"UGC": 'C',
	"UGA": '_',
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

var mapCodons = map[uint8][]string{
	'_': {"UAA", "UAG", "UGA"},
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

var mapProteinWeights = [...]int64{
	'A': 7103711,
	'C': 10300919,
	'D': 11502694,
	'E': 12904259,
	'F': 14706841,
	'G': 5702146,
	'H': 13705891,
	'I': 11308406,
	'K': 12809496,
	'L': 11308406,
	'M': 13104049,
	'N': 11404293,
	'P': 9705276,
	'Q': 12805858,
	'R': 15610111,
	'S': 8703203,
	'T': 10104768,
	'V': 9906841,
	'W': 18607931,
	'Y': 16306333,
}

func geneticStringToProtein(s []byte) ([]byte, int64, error) {
	if len(s)%3 != 0 {
		return nil, -1, fmt.Errorf("%w: length %d is not a multiple of 3", ErrInvalidGeneticString, len(s))
	}

	var weight int64
	data := make([]byte, 0, len(s)/3)

	for idx := 0; idx < len(s); idx += 3 {
		codon := string(s[idx : idx+3])

		aminoAcid := mapAminoAcids[codon]
		if aminoAcid == 0 {
			return nil, -1, fmt.Errorf("%w: codon %q is invalid", ErrInvalidGeneticString, codon)
		}

		if aminoAcid == '_' {
			break
		}

		data = append(data, aminoAcid)
		weight += mapProteinWeights[aminoAcid]
	}

	return data, weight, nil
}

type Protein struct {
	data   []byte
	weight int64
}

func NewProtein(b []byte) (*Protein, error) {
	protein := &Protein{
		data: b,
	}

	for _, char := range b {
		weight := mapProteinWeights[char]
		if weight == 0 {
			return nil, fmt.Errorf("%w: invalid amino acid %q", ErrInvalidGeneticString, char)
		}

		protein.weight += weight
	}

	return protein, nil
}

func NewProteinFromGeneticString(b []byte) (*Protein, error) {
	data, weight, err := geneticStringToProtein(b)
	if err != nil {
		return nil, err
	}

	return &Protein{
		data:   data,
		weight: weight,
	}, nil
}

func (p *Protein) String() string {
	return string(p.data)
}

func (p *Protein) Weight() int64 {
	return p.weight
}

func (p *Protein) WeightFloat() float64 {
	return float64(p.weight) / 100000
}

func (p *Protein) InferMRNA() int64 {
	if len(p.data) == 0 {
		return 0
	}

	count := int64(1)

	for _, char := range p.data {
		count *= int64(len(mapCodons[char]))
		count %= 1000000
	}

	return (count * 3) % 1000000
}
