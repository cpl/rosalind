package rosalind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblemDNA(t *testing.T) {
	t.Parallel()

	dna := "AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC"
	count := CountACGT(dna)

	assert.Equal(t, 20, count['A'])
	assert.Equal(t, 12, count['C'])
	assert.Equal(t, 17, count['G'])
	assert.Equal(t, 21, count['T'])
}

func TestProblemRNA(t *testing.T) {
	t.Parallel()

	dna := "GATGGAACTTGACTACGTAAATT"
	rna := DNAtoRNA(dna)

	assert.Equal(t, "GAUGGAACUUGACUACGUAAAUU", rna)
}
