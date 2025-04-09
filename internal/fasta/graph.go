package fasta

func Graph(faa *FASTA, prefixLen int) map[string][]string {
	m := make(map[string][]string, len(faa.Sequences))

	for idx := range faa.Sequences {
		prefix := string(faa.Sequences[idx].Data[:prefixLen])
		m[prefix] = append(m[prefix], faa.Sequences[idx].Label)
	}

	return m
}

func OverlapGraph(graph map[string][]string) [][2]string {
	out := make([][2]string, 0)
	for _, values := range graph {
		for idx, value := range values {
			for _, value2 := range values[idx+1:] {
				out = append(out, [2]string{value, value2})
			}
		}
	}

	return out
}
