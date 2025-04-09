package fasta

const (
	MIMEType      = "text/x-fasta"
	FileExtension = ".fasta"
)

var (
	fastaLabelIdentifier = []byte{'>'}
	fastaNewLine         = []byte{'\n'}
)
