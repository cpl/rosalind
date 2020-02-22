package rosalind

import (
	"bufio"
	"io"
	"strings"
)

func GCContent(dna string) float64 {
	gcCount := 0
	for _, symbol := range dna {
		switch symbol {
		case 'G', 'C':
			gcCount++
		}
	}

	return (float64(gcCount) / float64(len(dna))) * 100.0
}

type FASTA map[string]string

func (fasta FASTA) DNAs() []string {
	dnas := make([]string, len(fasta))

	index := 0
	for _, dna := range fasta {
		dnas[index] = dna
		index++
	}
	return dnas
}

func (fasta FASTA) MaxGCContent() (string, float64) {
	maxkey := ""
	maxGC := 0.0
	for key, dna := range fasta {
		if gc := GCContent(dna); gc > maxGC {
			maxGC = gc
			maxkey = key
		}
	}
	return maxkey, maxGC
}

func (fasta FASTA) GCContents() map[string]float64 {
	ret := make(map[string]float64, len(fasta))
	for key, dna := range fasta {
		ret[key] = GCContent(dna)
	}
	return ret
}

func LoadFASTA(in io.Reader) FASTA {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)

	fasta := make(FASTA)
	var key string

	for scanner.Scan() {
		val := scanner.Text()
		if strings.TrimSpace(val) == "" {
			continue
		}
		if val[0] == '>' {
			key = val[1:]
			val = ""
		}

		fasta[key] = fasta[key] + val
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return fasta
}
