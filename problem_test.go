package rosalind

import (
	"bytes"
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

func TestProblemREVC(t *testing.T) {
	t.Parallel()

	dna := "AAAACCCGGT"
	revc := ReverseComplement(dna)

	assert.Equal(t, "ACCGGGTTTT", revc)
}

func TestProblemFIB(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 19, Rabbits(5, 3))
}

func TestProblemGC(t *testing.T) {
	t.Parallel()

	gcPercent := GCContent("AGCTATAG")
	assert.Equal(t, 37.5, gcPercent)

	fastaData := `
>Rosalind_6404
CCTGCGGAAGATCGGCACTAGAATAGCCAGAACCGTTTCTCTGAGGCTTCCGGCCTTCCC
TCCCACTAATAATTCTGAGG
>Rosalind_5959
CCATCGGTAGCGCATCCTTAGTCCAATTAAGTCCCTATCCAGGCGCTCCGCCGAAGGTCT
ATATCCATTTGTCAGCAGACACGC
>Rosalind_0808
CCACCCTCGTGGTATGGCTAGGCATTCAGGAACCGGAGAACGCTTCAGACCAGCCCGGAC
TGGGAACCTGCGGGCAGTAGGTGGAAT
`

	buffer := bytes.NewBufferString(fastaData)
	fasta := LoadFASTA(buffer)

	maxkey, maxGC := fasta.MaxGCContent()

	assert.Equal(t, "Rosalind_0808", maxkey)
	assert.Equal(t, 60.91954022988506, maxGC)
}

func TestProblemHAMM(t *testing.T) {
	t.Parallel()
	assert.Equal(t, 7, HammingDistance("GAGCCTACTAACGGGAT", "CATCGTAATGACGGCCT"))
}

func TestProblemPROT(t *testing.T) {
	t.Parallel()
	rna := "AUGGCCAUGGCGCCCAGAACUGAGAUCAAUAGUACCCGUAUUAACGGGUGA"
	assert.Equal(t, "MAMAPRTEINSTRING", RNAtoProtein(rna))
}
