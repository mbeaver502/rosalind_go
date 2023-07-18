package nucleotide

import "fmt"

// Sequence represents a sequence of zero or more Nucleotides.
type Sequence []Nucleotide

// Convert the given string s into a slice of Nucleotide.
func ToSequence(s string) Sequence {
	nts := Sequence{}
	for _, n := range s {
		nt := Nucleotide(n)
		switch nt {
		case Adenine, Cytosine, Guanine, Thymine, Uracil:
			nts = append(nts, nt)
		}
	}
	return nts
}

// String returns the string representation of the NucleotideSequence.
func (ns Sequence) String() string {
	return fmt.Sprintf("%s\n", string(ns))
}
