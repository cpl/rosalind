package rosalind

import (
	"reflect"
	"strings"
	"testing"
)

func Test_Motif(t *testing.T) {
	t.Parallel()

	tests := map[string][]int{
		"GATATATGCATATACTT/ATAT":  {2, 4, 10},
		"ATATAT/ATAT":             {1, 3},
		"ATGAGA/ATAT":             {},
		"AUGCUUCAGAAAGGUCUUACG/U": {2, 5, 6, 15, 17, 18},
	}

	for test, want := range tests {
		t.Run(test, func(t *testing.T) {
			split := strings.Split(test, "/")
			data := []byte(split[0])
			motif := []byte(split[1])

			motifs := Motif(data, motif)

			if !reflect.DeepEqual(motifs, want) {
				t.Errorf("expected %v, got %v", want, motifs)
			}
		})
	}
}

func Test_HammingDistance(t *testing.T) {
	t.Parallel()

	tests := map[string]struct{ num int }{
		"GAGCCTACTAACGGGAT/CATCGTAATGACGGCCT": {num: 7},
		"CATCGTAATGACGGCCT/CATCGTAATGACGGCCT": {num: 0},
	}

	for test, want := range tests {
		t.Run(test, func(t *testing.T) {
			split := strings.Split(test, "/")
			a := []byte(split[0])
			b := []byte(split[1])

			distance, err := HammingDistance(a, b)
			if err != nil {
				t.Fatal(err)
			}

			if distance != want.num {
				t.Errorf("expected %d, got %d", want.num, distance)
			}
		})
	}
}
