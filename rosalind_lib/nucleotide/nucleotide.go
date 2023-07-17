package nucleotide

import (
	"fmt"
	"strings"
)

// Nucleotide represents a single nucleotide in a DNA sequence.
type Nucleotide rune

// NucleotideSequence represents a sequence of zero or more Nucleotides.
type NucleotideSequence []Nucleotide

// NucleotideCounts maps the number of appearances of Nucleotides.
// Example: counts['A'] = 5
type NucleotideCounts map[Nucleotide]uint

// NucleotideMapping maps one Nucleotide to another.
// Example: 'A' -> 'T'
type NucleotideMapping map[Nucleotide]Nucleotide

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

// String returns the string representation of the NucleotideSequence.
func (ns NucleotideSequence) String() string {
	return string(ns)
}

// String returns the string representation of the Nucleotide.
func (nt Nucleotide) String() string {
	return string(nt)
}

// String returns a pretty printed representation of the NucleotideCounts.
// Currently only returns DNA nucleotides.
func (nc NucleotideCounts) String() string {
	return fmt.Sprintf(
		"%d %d %d %d",
		nc[Adenine],
		nc[Cytosine],
		nc[Guanine],
		nc[Thymine],
	)
}

// String returns a pretty printed representation of the NucleotideMapping.
func (nm NucleotideMapping) String() string {
	var sb strings.Builder
	for k, v := range nm {
		sb.WriteString(fmt.Sprintf("%s -> %s, ", k, v))
	}
	return fmt.Sprintf("{%s}", strings.TrimRight(sb.String(), ", "))
}
