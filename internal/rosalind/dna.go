package rosalind

import (
	"bytes"
	"fmt"
	"sync"
)

type DNA struct {
	data    []byte
	counted bool
	a       int
	c       int
	g       int
	t       int
	gc      float64
}

func NewDNA(b []byte) (*DNA, error) {
	for idx, char := range b {
		if char >= 'a' && char <= 'z' {
			char -= 'a' - 'A'
			b[idx] = char
		}

		switch char {
		case 'A', 'C', 'G', 'T':
		default:
			//panic("invalid DNA character")
			return nil, fmt.Errorf("%w: invalid DNA character '%c'", ErrInvalidGeneticString, char)
		}
	}

	return &DNA{data: b, gc: -1}, nil
}

func NewDNAString(s string) (*DNA, error) {
	return NewDNA([]byte(s))
}

func (dna *DNA) Count() (a, c, g, t int) {
	if dna.counted {
		return dna.a, dna.c, dna.g, dna.t
	}

	dna.a = bytes.Count(dna.data, []byte("A"))
	dna.c = bytes.Count(dna.data, []byte("C"))
	dna.g = bytes.Count(dna.data, []byte("G"))
	dna.t = bytes.Count(dna.data, []byte("T"))
	dna.counted = true

	return dna.a, dna.c, dna.g, dna.t
}

func (dna *DNA) CountThreaded() (a, c, g, t int) {
	if dna.counted {
		return dna.a, dna.c, dna.g, dna.t
	}

	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		dna.a = bytes.Count(dna.data, []byte("A"))
		wg.Done()
	}()

	go func() {
		dna.c = bytes.Count(dna.data, []byte("C"))
		wg.Done()
	}()

	go func() {
		dna.g = bytes.Count(dna.data, []byte("G"))
		wg.Done()
	}()

	go func() {
		dna.t = bytes.Count(dna.data, []byte("T"))
		wg.Done()
	}()

	wg.Wait()
	dna.counted = true

	return dna.a, dna.c, dna.g, dna.t
}

func (dna *DNA) String() string {
	return string(dna.data)
}

func (dna *DNA) StringReverse() string {
	dataLen := len(dna.data)
	out := make([]byte, dataLen)

	for idx := 0; idx < dataLen; idx++ {
		out[idx] = dna.data[dataLen-idx-1]
	}

	return string(out)
}

func (dna *DNA) ToRNA() *RNA {
	return &RNA{
		data:    bytes.ReplaceAll(dna.data, []byte("T"), []byte("U")),
		counted: dna.counted,
		a:       dna.a,
		c:       dna.c,
		g:       dna.g,
		u:       dna.t,
	}
}

func (dna *DNA) Complement() *DNA {
	data := make([]byte, len(dna.data))
	for idx, char := range dna.data {
		switch char {
		case 'A':
			data[idx] = 'T'
		case 'C':
			data[idx] = 'G'
		case 'G':
			data[idx] = 'C'
		case 'T':
			data[idx] = 'A'
		}
	}

	return &DNA{
		data:    data,
		counted: dna.counted,
		a:       dna.t,
		c:       dna.g,
		g:       dna.c,
		t:       dna.a,
	}
}

func (dna *DNA) GCContentPercent() float64 {
	if dna.gc != -1 {
		return dna.gc
	}

	dnaLen := len(dna.data)
	if dnaLen == 0 {
		return 0
	}

	if !dna.counted {
		dna.Count()
	}

	gcCount := dna.c + dna.g
	gcPercent := float64(gcCount) / float64(dnaLen)
	dna.gc = gcPercent

	return dna.gc
}

func (dna *DNA) ToProtein() (*Protein, error) {
	proteinString, weight, err := geneticStringToProtein(dna.data)
	if err != nil {
		return nil, fmt.Errorf("error converting DNA to protein: %w", err)
	}

	return &Protein{
		data:   proteinString,
		weight: weight,
	}, nil
}
