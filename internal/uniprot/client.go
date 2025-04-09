package uniprot

import (
	"net/http"
	"time"

	"github.com/cpl/rosalind/internal/fasta"
)

var httpClient = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: nil,
	},
	Timeout: 15 * time.Second,
}

type Client interface {
	GetFASTA(id string) (*fasta.FASTA, error)
}
