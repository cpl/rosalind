package rosalind

import "fmt"

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
