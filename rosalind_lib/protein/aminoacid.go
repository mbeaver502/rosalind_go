package protein

import "regexp"

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

var monoisotopicMass = map[AminoAcid]float64{
	A: 71.03711,
	C: 103.00919,
	D: 115.02694,
	E: 129.04259,
	F: 147.06841,
	G: 57.02146,
	H: 137.05891,
	I: 113.08406,
	K: 128.09496,
	L: 113.08406,
	M: 131.04049,
	N: 114.04293,
	P: 97.05276,
	Q: 128.05858,
	R: 156.10111,
	S: 87.03203,
	T: 101.04768,
	V: 99.06841,
	W: 186.07931,
	Y: 163.06333,
}

func isValidAminoAcid(s string) bool {
	rex := regexp.MustCompile(`[ACDEFGHIKLMNPQRSTVWY]+`)
	return rex.MatchString(s)
}

// String converts the given AminoAcid into a string.
func (a AminoAcid) String() string {
	return string(a)
}
