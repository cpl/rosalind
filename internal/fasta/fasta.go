package fasta

import (
	"bufio"
	"io"
)

type Sequence struct {
	Label string
	Data  []byte
}

type FASTA struct {
	lookup    map[string]int
	Sequences []Sequence
}

func Parse(r io.Reader) *FASTA {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	fasta := &FASTA{
		Sequences: make([]Sequence, 0, 16),
	}

	var sequence Sequence

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
