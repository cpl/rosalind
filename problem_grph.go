package rosalind

import "strings"

func (fasta FASTA) Graph(level int) [][]string {
	var ret [][]string

	for key0, value0 := range fasta {
		for key1, value1 := range fasta {
			if key0 == key1 {
				continue
			}

			if strings.HasSuffix(value0, value1[:level]) {
				ret = append(ret, []string{key0, key1})
			}
		}
	}

	return ret
}
