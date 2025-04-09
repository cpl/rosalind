package uniprot

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cpl/rosalind/internal/fasta"
)

var DefaultCacheClient = &CacheClient{
	Dir: filepath.Join("/tmp", "/uniprot_cache"),
}

type CacheClient struct {
	Dir string
}

func (client *CacheClient) GetFASTA(id string) (*fasta.FASTA, error) {
	_ = os.MkdirAll(client.Dir, 0o700)

	fp, err := os.OpenFile(filepath.Join(client.Dir, id+fasta.FileExtension), os.O_RDWR|os.O_CREATE, 0o600)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	if stat, _ := fp.Stat(); stat.Size() > 0 {
		return fasta.Parse(fp), nil
	}

	req, err := http.NewRequest("GET", "https://rest.uniprot.org/uniprotkb/"+id+".fasta", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("User-Agent", "dev@cpl.li")

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("error response status: %d", res.StatusCode)
	}

	r, err := gzip.NewReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error creating gzip reader: %w", err)
	}
	defer r.Close()

	faa := fasta.Parse(r)
	if _, err = faa.WriteTo(fp); err != nil {
		return nil, fmt.Errorf("error writing fasta to cache: %w", err)
	}

	return faa, nil
}
