package nucleotide

// Nucleotide represents a single nucleotide in a DNA sequence.
type Nucleotide rune

const (
	Adenine  = Nucleotide('A')
	Cytosine = Nucleotide('C')
	Guanine  = Nucleotide('G')
	Thymine  = Nucleotide('T')
	Uracil   = Nucleotide('U')
)

// Convert the given string s into a slice of Nucleotide.
func ToSlice(s string) []Nucleotide {
	// len(s) will give us the number of bytes
	// instead of the number of runes, but
	// since our runes are in ASCII, this is okay.
	nts := make([]Nucleotide, len(s))
	for i, n := range s {
		nt := Nucleotide(n)
		switch nt {
		case Adenine, Cytosine, Guanine, Thymine:
			nts[i] = nt
		}
	}
	return nts
}
