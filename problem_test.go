package rosalind

import (
	"bytes"
	"fmt"
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

func TestProblemSUBS(t *testing.T) {
	t.Parallel()

	dna := "GATATATGCATATACTT"
	assert.Equal(t, []int{2, 4, 10}, Motif(dna, "ATAT"))
}

func TestProblemMRNA(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 12, ProteinToRNACount("MA\x00"))
}

func TestProblemPRTM(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 821.3919199999999, ProteinMass("SKADYEK"))
}

func TestProblemCONS(t *testing.T) {
	t.Parallel()

	fastaData := `
>Rosalind_1
ATCCAGCT
>Rosalind_2
GGGCAACT
>Rosalind_3
ATGGATCT
>Rosalind_4
AAGCAACC
>Rosalind_5
TTGGAACT
>Rosalind_6
ATGCCATT
>Rosalind_7
ATGGCACT
`

	buffer := bytes.NewBufferString(fastaData)
	fasta := LoadFASTA(buffer)

	matrix := DNAtoProfileMatrix(fasta.DNAs()...)
	assert.Equal(t, "ATGCAACT", matrix.ConsensusString())
}

func TestProblemMPRT(t *testing.T) {
	t.Parallel()

	ids := map[string][]int{
		"A2Z669":            {},
		"B5ZC00":            {85, 118, 142, 306, 395},
		"P07204_TRBM_HUMAN": {47, 115, 116, 382, 409},
		"P20840_SAG1_YEAST": {79, 109, 135, 248, 306, 348, 364, 402, 485, 501, 614},
	}

	for id, expected := range ids {
		protein, err := GetUniProt(id)
		assert.NoError(t, err, id)
		fmt.Println(id)
		loc := Motif2(protein, "N{P}[ST]{P}")
		assert.Equal(t, expected, loc, id)
	}
}

func TestProblemGRPH(t *testing.T) {
	t.Parallel()

	fastaData := `
>Rosalind_0498
AAATAAA
>Rosalind_2391
AAATTTT
>Rosalind_2323
TTTTCCC
>Rosalind_0442
AAATCCC
>Rosalind_5013
GGGTGGG
`

	buffer := bytes.NewBufferString(fastaData)
	fasta := LoadFASTA(buffer)

	assert.Equal(t, [][]string{
		{"Rosalind_0498", "Rosalind_2391"},
		{"Rosalind_0498", "Rosalind_0442"},
		{"Rosalind_2391", "Rosalind_2323"},
	}, fasta.Graph(3))
}

func TestProblemORF(t *testing.T) {
	t.Parallel()

	dna := "AGCCATGTAGCTAACTCAGGTTACATGGGGATGACCCCGCGACTTGGATTAGAGTCTCTTTTGGAATAAGCCTGAATGATCCGAGTAGCATCTCAG"
	frames := ORF(dna)
	framesDistinct := make(map[string]int)

	for _, frame := range frames {
		framesDistinct[RNAtoProtein(frame)] = 0
	}

	assert.Equal(t, map[string]int{
		"MLLGSFRLIPKETLIQVAGSSPCNLS": 0,
		"M":                          0,
		"MGMTPRLGLESLLE":             0,
		"MTPRLGLESLLE":               0,
	}, framesDistinct)
}
