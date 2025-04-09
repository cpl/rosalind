package uniprot_test

import (
	"testing"

	"github.com/cpl/rosalind/internal/uniprot"
)

func TestUniprotDownloadSequence(t *testing.T) {
	t.Parallel()

	fasta, err := uniprot.DefaultCacheClient.GetFASTA("B5ZC00")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(fasta.Sequences)
}
