package protein

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mbeaver502/rosalind_go/rosalind_lib/nucleotide"
)

// Protein represents a sequence of AminoAcid.
type Protein []AminoAcid

var translationMap = map[string]AminoAcid{
	"UUU": F, "CUU": L, "AUU": I, "GUU": V,
	"UUC": F, "CUC": L, "AUC": I, "GUC": V,
	"UUA": L, "CUA": L, "AUA": I, "GUA": V,
	"UUG": L, "CUG": L, "AUG": M, "GUG": V,
	"UCU": S, "CCU": P, "ACU": T, "GCU": A,
	"UCC": S, "CCC": P, "ACC": T, "GCC": A,
	"UCA": S, "CCA": P, "ACA": T, "GCA": A,
	"UCG": S, "CCG": P, "ACG": T, "GCG": A,
	"UAU": Y, "CAU": H, "AAU": N, "GAU": D,
	"UAC": Y, "CAC": H, "AAC": N, "GAC": D,
	"CAA": Q, "AAA": K, "GAA": E,
	"CAG": Q, "AAG": K, "GAG": E,
	"UGU": C, "CGU": R, "AGU": S, "GGU": G,
	"UGC": C, "CGC": R, "AGC": S, "GGC": G,
	"CGA": R, "AGA": R, "GGA": G,
	"UGG": W, "CGG": R, "AGG": R, "GGG": G,
}

// New creates a new Protein from the given string.
// The string should contain valid amino acids only.
func New(s string) (Protein, error) {
	if !isValidAminoAcid(s) {
		return Protein{}, errors.New("protein: invalid input")
	}

	p := Protein{}
	for _, a := range s {
		p = append(p, AminoAcid(a))
	}
	return p, nil
}

// Translate will translate the given Sequence into Protein.
func Translate(ns nucleotide.Sequence) Protein {
	out := Protein{}

	for i := 0; i <= len(ns)-3; i += 3 {
		codon := fmt.Sprintf("%s%s%s", ns[i], ns[i+1], ns[i+2])
		if codon == "UAA" || codon == "UAG" || codon == "UGA" {
			break
		}
		out = append(out, translationMap[codon])
	}
	return out
}

// String converts the Protein p into a string.
func (p Protein) String() string {
	var sb strings.Builder
	for _, a := range p {
		sb.WriteRune(rune(a))
	}
	return fmt.Sprintf("%s\n", sb.String())
}

// Mass calculates monoisotopic mass of the Protein p.
func (p Protein) Mass() float64 {
	mass := float64(0)
	for _, a := range p {
		mass += monoisotopicMass[a]
	}
	return mass
}
