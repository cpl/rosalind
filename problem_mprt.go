package rosalind

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

const uniprotEndpoint = "https://www.uniprot.org/uniprot/"

func GetUniProt(id string) (string, error) {
	res, err := http.Get(uniprotEndpoint + id + ".fasta")
	if err != nil {
		return "", fmt.Errorf("failed get fasta uniprot, %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("bad status code, %d", res.StatusCode)
	}

	dnas := LoadFASTA(res.Body).DNAs()
	if len(dnas) == 0 {
		return "", nil
	}

	return dnas[0], nil
}

func Motif2(str, search string) []int {
	search = strings.ReplaceAll(search, "{", "[^")
	search = strings.ReplaceAll(search, "}", "]")

	re, err := regexp.Compile(search)
	if err != nil {
		panic(err)
	}

	var ret []int

	index := 0
	for {
		loc := re.FindStringIndex(str[index:])
		if loc == nil || len(loc) == 0 {
			if index == 0 {
				return []int{}
			}
			return ret
		}
		index = loc[0] + 1 + index
		ret = append(ret, index)
	}
}
