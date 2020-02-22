package rosalind

import "fmt"

type ProfileMatrix map[rune][]int

func (matrix ProfileMatrix) ConsensusString() string {
	l := len(matrix['A'])
	ret := make([]rune, l)

	for index := 0; index < l; index++ {
		v := 0

		a := matrix['A'][index]
		if a > v {
			v = a
			ret[index] = 'A'
		}

		c := matrix['C'][index]
		if c > v {
			v = c
			ret[index] = 'C'
		}

		g := matrix['G'][index]
		if g > v {
			v = g
			ret[index] = 'G'
		}

		t := matrix['T'][index]
		if t > v {
			v = t
			ret[index] = 'T'
		}
	}

	return string(ret)
}

func (matrix ProfileMatrix) String() string {
	return fmt.Sprintf("A: %v\nC: %v\nG: %v\nT: %v\n",
		matrix['A'], matrix['C'], matrix['G'], matrix['T'])
}

func DNAtoProfileMatrix(dnas ...string) ProfileMatrix {
	if len(dnas) == 0 {
		return nil
	}

	matrix := ProfileMatrix{
		'A': make([]int, len(dnas[0])),
		'C': make([]int, len(dnas[0])),
		'G': make([]int, len(dnas[0])),
		'T': make([]int, len(dnas[0])),
	}

	for _, dna := range dnas {
		for index, symbol := range dna {
			matrix[symbol][index] += 1
		}
	}

	return matrix
}
