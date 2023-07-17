package nucleotide

// Nucleotide represents a single nucleotide in a DNA sequence.
type Nucleotide rune

// NucleotideSequence represents a sequence of zero or more Nucleotides.
type NucleotideSequence []Nucleotide

const (
	Adenine  = Nucleotide('A')
	Cytosine = Nucleotide('C')
	Guanine  = Nucleotide('G')
	Thymine  = Nucleotide('T')
	Uracil   = Nucleotide('U')
)

// Convert the given string s into a slice of Nucleotide.
func ToSlice(s string) NucleotideSequence {
	nts := NucleotideSequence{}
	for _, n := range s {
		nt := Nucleotide(n)
		switch nt {
		case Adenine, Cytosine, Guanine, Thymine, Uracil:
			nts = append(nts, nt)
		}
	}
	return nts
}
