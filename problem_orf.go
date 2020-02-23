package rosalind

const (
	StartCodonDNA = "ATG"
	StartCodonRNA = "AUG"
)

func ORF(dna string) (ret []string) {
	rna := DNAtoRNA(dna)

	ret = append(ret, orf(rna)...)
	ret = append(ret, orf(rna[1:len(rna)-2])...)
	ret = append(ret, orf(rna[2:len(rna)-1])...)

	rna = DNAtoRNA(ReverseComplement(dna))

	ret = append(ret, orf(rna)...)
	ret = append(ret, orf(rna[1:len(rna)-2])...)
	ret = append(ret, orf(rna[2:len(rna)-1])...)

	return ret
}

func orf(rna string) (ret []string) {
	var (
		building bool
		starts   []int
	)

	for index := 0; index < len(rna)-3; index += 3 {
		str := rna[index : index+3]

		if str == StartCodonRNA {
			starts = append(starts, index)
			building = true
		}
		if building {
			for _, end := range CodonRNATable[0] {
				if str == end {
					for _, start := range starts {
						ret = append(ret, rna[start:index+3])
					}
					starts = []int{}
					building = false
					break
				}
			}
		}
	}

	return ret
}
