package rosalind

import (
	"bufio"
	"io"
)

type FASTA struct {
	Labels []string
	Data   [][]byte
}

func ParseFASTA(r io.Reader) *FASTA {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	labels := make([]string, 0)
	data := make([][]byte, 0)
	buffer := make([]byte, 0, 1<<10)

	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		if line[0] == '>' {
			labels = append(labels, string(line[1:]))
			if len(labels) > 1 {
				data = append(data, buffer)
				buffer = make([]byte, 0, 1<<10)
			}

			continue
		}

		buffer = append(buffer, line...)
	}

	data = append(data, buffer)

	return &FASTA{
		Labels: labels,
		Data:   data,
	}
}

func (fasta *FASTA) GetData(label string) []byte {
	for idx, l := range fasta.Labels {
		if l == label {
			return fasta.Data[idx]
		}
	}

	return nil
}

func (fasta *FASTA) WriteTo(w io.Writer) (n int64, err error) {
	var wN int

	for idx, label := range fasta.Labels {
		data := fasta.Data[idx]

		wN, err = fasta.WriteToLabel(w, label, data, 60)
		if err != nil {
			return n + int64(wN), err
		}

		n += int64(wN)
	}

	return n, nil
}

func (fasta *FASTA) WriteToLabel(w io.Writer, label string, data []byte, lineLen int) (n int, err error) {
	var wN int

	wN, err = w.Write([]byte(">" + label + "\n"))
	if err != nil {
		return n + wN, err
	}
	n += wN

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
			wN, err = w.Write([]byte("\n"))
			if err != nil {
				return n + wN, err
			}
			n += wN
		}
	}

	wN, err = w.Write([]byte("\n"))
	if err != nil {
		return n + wN, err
	}

	return n + wN, nil
}

func (fasta *FASTA) Graph(prefixLen int) map[string][]string {
	m := make(map[string][]string)

	for idx, data := range fasta.Data {
		prefix := string(data[:prefixLen])
		m[prefix] = append(m[prefix], fasta.Labels[idx])
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
