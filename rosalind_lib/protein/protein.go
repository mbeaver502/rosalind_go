package protein

import (
	"fmt"
	"strings"

	"github.com/mbeaver502/rosalind_go/rosalind_lib/nucleotide"
)

type AminoAcid rune

type Protein []AminoAcid

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

var translationMap = map[string]AminoAcid{
	"UUU": 'F', "CUU": 'L', "AUU": 'I', "GUU": 'V',
	"UUC": 'F', "CUC": 'L', "AUC": 'I', "GUC": 'V',
	"UUA": 'L', "CUA": 'L', "AUA": 'I', "GUA": 'V',
	"UUG": 'L', "CUG": 'L', "AUG": 'M', "GUG": 'V',
	"UCU": 'S', "CCU": 'P', "ACU": 'T', "GCU": 'A',
	"UCC": 'S', "CCC": 'P', "ACC": 'T', "GCC": 'A',
	"UCA": 'S', "CCA": 'P', "ACA": 'T', "GCA": 'A',
	"UCG": 'S', "CCG": 'P', "ACG": 'T', "GCG": 'A',
	"UAU": 'Y', "CAU": 'H', "AAU": 'N', "GAU": 'D',
	"UAC": 'Y', "CAC": 'H', "AAC": 'N', "GAC": 'D',
	"CAA": 'Q', "AAA": 'K', "GAA": 'E',
	"CAG": 'Q', "AAG": 'K', "GAG": 'E',
	"UGU": 'C', "CGU": 'R', "AGU": 'S', "GGU": 'G',
	"UGC": 'C', "CGC": 'R', "AGC": 'S', "GGC": 'G',
	"CGA": 'R', "AGA": 'R', "GGA": 'G',
	"UGG": 'W', "CGG": 'R', "AGG": 'R', "GGG": 'G',
}

func Translate(ns nucleotide.Sequence) Protein {
	out := Protein{}

	for i := 0; i < len(ns)-3; i += 3 {
		codon := fmt.Sprintf("%s%s%s", string(ns[i]), string(ns[i+1]), string(ns[i+2]))
		if codon == "UAA" || codon == "UAG" || codon == "UGA" {
			break
		}
		out = append(out, translationMap[codon])
	}
	return out
}

func (p Protein) String() string {
	var sb strings.Builder
	for _, a := range p {
		sb.WriteRune(rune(a))
	}
	return fmt.Sprintf("%s\n", sb.String())
}
