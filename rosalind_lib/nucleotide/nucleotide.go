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

// String returns the string representation of the Nucleotide.
func (nt Nucleotide) String() string {
	return string(nt)
}
