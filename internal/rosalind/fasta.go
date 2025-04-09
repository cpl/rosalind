package rosalind

import (
	"bufio"
	"io"
)

type FASTASequence struct {
	Label string
	Data  []byte
}

type FASTA struct {
	lookup    map[string]int
	Sequences []FASTASequence
}

func ParseFASTA(r io.Reader) *FASTA {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	fasta := &FASTA{
		Sequences: make([]FASTASequence, 0, 16),
	}

	var sequence FASTASequence

	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		if line[0] != '>' {
			sequence.Data = append(sequence.Data, line...)
			continue
		}

		if sequence.Label != "" {
			fasta.Sequences = append(fasta.Sequences, sequence)
		}

		sequence.Label = string(line[1:])
		sequence.Data = make([]byte, 0, 1024)
	}

	if sequence.Label != "" {
		fasta.Sequences = append(fasta.Sequences, sequence)
	}

	if len(fasta.Sequences) > 16 {
		fasta.lookup = make(map[string]int, len(fasta.Sequences))
		for idx, seq := range fasta.Sequences {
			fasta.lookup[seq.Label] = idx
		}
	}

	return fasta
}

func (fasta *FASTA) IndexLabel(label string) int {
	if fasta.lookup != nil {
		return fasta.lookup[label]
	}

	for idx := range fasta.Sequences {
		if fasta.Sequences[idx].Label == label {
			return idx
		}
	}

	return -1
}

func (fasta *FASTA) GetData(label string) []byte {
	idx := fasta.IndexLabel(label)
	if idx == -1 {
		return nil
	}

	return fasta.Sequences[idx].Data
}

func (fasta *FASTA) WriteTo(w io.Writer) (n int64, err error) {
	var wN int

	for idx := range fasta.Sequences {
		data := fasta.Sequences[idx].Data

		wN, err = fasta.WriteToLabel(w, fasta.Sequences[idx].Label, data, 60)
		if err != nil {
			return n + int64(wN), err
		}

		n += int64(wN)
	}

	return n, nil
}

var (
	fastaLabelIdentifier = []byte{'>'}
	fastaNewLine         = []byte{'\n'}
)

func (fasta *FASTA) WriteToLabel(w io.Writer, label string, data []byte, lineLen int) (n int, err error) {
	var wN int

	_, _ = w.Write(fastaLabelIdentifier)
	wN, err = w.Write([]byte(label))
	if err != nil {
		return n + wN, err
	}
	_, _ = w.Write(fastaNewLine)

	n += wN + 2

	for len(data) > 0 {
		to := lineLen
		if to > len(data) {
			to = len(data)
		}

		wN, err = w.Write(data[:to])
		if err != nil {
			return n + wN, err
		}
		n += wN

		data = data[to:]

		if len(data) > 0 {
			wN, err = w.Write(fastaNewLine)
			if err != nil {
				return n + wN, err
			}
			n += wN
		}
	}

	wN, err = w.Write(fastaNewLine)
	if err != nil {
		return n + wN, err
	}

	return n + wN, nil
}

func (fasta *FASTA) Graph(prefixLen int) map[string][]string {
	m := make(map[string][]string, len(fasta.Sequences))

	for idx := range fasta.Sequences {
		prefix := string(fasta.Sequences[idx].Data[:prefixLen])
		m[prefix] = append(m[prefix], fasta.Sequences[idx].Label)
	}

	return m
}

func FASTAOverlapGraph(graph map[string][]string) [][2]string {
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
