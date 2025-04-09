package rosalind

var (
	bytesDNAA = []byte{'A'}
	bytesDNAC = []byte{'C'}
	bytesDNAG = []byte{'G'}
	bytesDNAT = []byte{'T'}

	bytesRNAU = []byte{'U'}
)

var dnaComplement = [0x100]byte{
	'A': 'T',
	'C': 'G',
	'G': 'C',
	'T': 'A',
}
