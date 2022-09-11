package rosalind

import "fmt"

type RNA struct {
	data    []byte
	counted bool
	a       int
	c       int
	g       int
	u       int
}

func (rna *RNA) String() string {
	return string(rna.data)
}

func (rna *RNA) ToProtein() (*Protein, error) {
	proteinString, weight, err := geneticStringToProtein(rna.data)
	if err != nil {
		return nil, fmt.Errorf("error converting RNA to protein: %w", err)
	}

	return &Protein{
		data:   proteinString,
		weight: weight,
	}, nil
}
