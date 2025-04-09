package rosalind

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/cpl/rosalind/internal/fasta"
)

var defaultUniprotHttpClient = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: nil,
	},
	Timeout: 30 * time.Second,
}

type httpDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

func UniprotDownloadSequence(httpClient httpDoer, id string) (*fasta.FASTA, error) {
	path := filepath.Join("./", "data", "uniprot", id+".fasta")
	pathDir := filepath.Dir(path)
	_ = os.MkdirAll(pathDir, 0o755)

	fp, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0o644)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", id, err)
	}
	defer fp.Close()

	if stat, _ := fp.Stat(); stat.Size() > 0 {
		return fasta.Parse(fp), nil
	}

	req, err := http.NewRequest("GET", "https://rest.uniprot.org/uniprotkb/"+id+".fasta", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error response status: %d", res.StatusCode)
	}

	r, _ := gzip.NewReader(res.Body)
	defer r.Close()
	data, _ := io.ReadAll(r)

	fasta := fasta.Parse(bytes.NewReader(data))
	_, err = fasta.WriteTo(fp)
	if err != nil {
		return nil, fmt.Errorf("error writing fasta to file: %w", err)
	}

	return fasta, nil
}
