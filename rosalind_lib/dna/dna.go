package dna

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/mbeaver502/rosalind_go/rosalind_lib/nucleotide"
)

// Dna represents a strand of DNA.
type Dna struct {
	Dna string
	dna []nucleotide.Nucleotide
}

// ErrInvalidDnaInput indicates that the provided string
// represents an invalid sequence of characters and
// cannot constitute a valid DNA sequence.
var ErrInvalidDnaInput = errors.New("dna: invalid input string")

// New creates a new Dna instance and returns a pointer to it.
func New(s string) (*Dna, error) {
	if !isValidString(s) {
		return nil, ErrInvalidDnaInput
	}

	su := strings.ToUpper(s)

	return &Dna{
		Dna: su,
		dna: toNucleotideSlice(su),
	}, nil
}

// Determine whether the given string contains only
// valid character representations of DNA nucleotides.
func isValidString(s string) bool {
	rex := regexp.MustCompile(`^[aAcCgGtT]+$`)
	return rex.MatchString(s)
}

// Convert the given string s into a slice of Nucleotide.
func toNucleotideSlice(s string) []nucleotide.Nucleotide {
	// len(s) will give us the number of bytes
	// instead of the number of runes, but
	// since our runes are in ASCII, this is okay.
	nts := make([]nucleotide.Nucleotide, len(s))
	for i, n := range s {
		nt := nucleotide.Nucleotide(n)
		switch nt {
		case nucleotide.Adenine, nucleotide.Cytosine, nucleotide.Guanine, nucleotide.Thymine:
			nts[i] = nt
		}
	}
	return nts
}

// String returns a string version of the Dna instance.
func (d *Dna) String() string {
	return d.Dna
}

// CountNucleotides counts the number of nucleotides in the Dna instance.
func (d *Dna) CountNucleotides() map[nucleotide.Nucleotide]uint {
	return d.countNucleotides()
}

// CountNucleotidesString counts the number of nucleotides in the Dna instance
// and returns the result in a pretty-printed string.
func (d *Dna) CountNucleotidesString() string {
	nt := d.countNucleotides()
	return fmt.Sprintf(
		"%d %d %d %d",
		nt[nucleotide.Adenine],
		nt[nucleotide.Cytosine],
		nt[nucleotide.Guanine],
		nt[nucleotide.Thymine],
	)
}

// Count the number of nucleotides in the Dna instance.
// If the given instance is nil, all counts default to
// zero values.
func (d *Dna) countNucleotides() map[nucleotide.Nucleotide]uint {
	nt := make(map[nucleotide.Nucleotide]uint)
	if d == nil {
		return nt
	}

	for _, n := range d.Dna {
		m := nucleotide.Nucleotide(n)
		switch m {
		case nucleotide.Adenine, nucleotide.Cytosine, nucleotide.Guanine, nucleotide.Thymine:
			nt[m]++
		}
	}
	return nt
}
