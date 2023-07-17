package dna

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Dna represents a strand of DNA.
type Dna struct {
	Dna string
}

// Nucleotide represents a single nucleotide in a DNA sequence.
type Nucleotide rune

const (
	NucleotideA = Nucleotide('A')
	NucleotideC = Nucleotide('C')
	NucleotideG = Nucleotide('G')
	NucleotideT = Nucleotide('T')
)

// ErrInvalidDnaInput indicates that the provided string
// represents an invalid sequence of characters and
// cannot constitute a valid DNA sequence.
var ErrInvalidDnaInput = errors.New("dna: invalid input string")

// New creates a new Dna instance and returns a pointer to it.
func New(s string) (*Dna, error) {
	if !isValidString(s) {
		return nil, ErrInvalidDnaInput
	}

	return &Dna{
		Dna: strings.ToUpper(s),
	}, nil
}

// Determine whether the given string contains only
// valid character representations of DNA nucleotides.
func isValidString(s string) bool {
	rex := regexp.MustCompile(`^[aAcCgGtT]+$`)
	return rex.MatchString(s)
}

// String returns a string version of the Dna instance.
func (d *Dna) String() string {
	return d.Dna
}

// CountNucleotides counts the number of nucleotides in the Dna instance.
func (d *Dna) CountNucleotides() map[Nucleotide]uint {
	return d.countNucleotides()
}

// CountNucleotidesString counts the number of nucleotides in the Dna instance
// and returns the result in a pretty-printed string.
func (d *Dna) CountNucleotidesString() string {
	nt := d.countNucleotides()
	return fmt.Sprintf(
		"%d %d %d %d",
		nt[NucleotideA],
		nt[NucleotideC],
		nt[NucleotideG],
		nt[NucleotideT],
	)
}

// Count the number of nucleotides in the Dna instance.
// If the given instance is nil, all counts default to
// zero values.
func (d *Dna) countNucleotides() map[Nucleotide]uint {
	nt := make(map[Nucleotide]uint)
	if d == nil {
		return nt
	}

	for _, n := range d.Dna {
		m := Nucleotide(n)
		switch m {
		case NucleotideA, NucleotideC, NucleotideG, NucleotideT:
			nt[m]++
		}
	}
	return nt
}
