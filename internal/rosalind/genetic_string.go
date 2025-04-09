package rosalind

import (
	"bytes"
	"fmt"
)

func Motif(data, motif []byte) []int {
	if len(motif) == 0 {
		return []int{}
	}

	idxList := make([]int, 0, bytes.Count(data, motif))
	offset := 1

	for {
		idx := bytes.Index(data, motif)
		if idx == -1 {
			break
		}

		idxList = append(idxList, idx+offset)
		data = data[idx+1:]

		offset += idx + 1
	}

	return idxList
}

func HammingDistance(a, b []byte) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("%w: %d != %d", ErrInvalidInput, len(a), len(b))
	}

	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
