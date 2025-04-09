package rosalind

import (
	"testing"
)

func TestUniprotDownloadSequence(t *testing.T) {
	t.Parallel()

	fasta, err := UniprotDownloadSequence(defaultUniprotHttpClient, "B5ZC00")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(fasta.Sequences)
}
