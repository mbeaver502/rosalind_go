package protein

// AminoAcid represents a single abbreviated
// amino acid that may make up a Protein.
type AminoAcid rune

const (
	A = AminoAcid('A')
	C = AminoAcid('C')
	D = AminoAcid('D')
	E = AminoAcid('E')
	F = AminoAcid('F')
	G = AminoAcid('G')
	H = AminoAcid('H')
	I = AminoAcid('I')
	K = AminoAcid('K')
	L = AminoAcid('L')
	M = AminoAcid('M')
	N = AminoAcid('N')
	P = AminoAcid('P')
	Q = AminoAcid('Q')
	R = AminoAcid('R')
	S = AminoAcid('S')
	T = AminoAcid('T')
	V = AminoAcid('V')
	W = AminoAcid('W')
	Y = AminoAcid('Y')
)

// String converts the given AminoAcid into a string.
func (a AminoAcid) String() string {
	return string(a)
}
