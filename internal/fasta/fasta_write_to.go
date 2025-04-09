package fasta

import (
	"io"
)

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
