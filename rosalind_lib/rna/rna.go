package rna

import (
	"errors"
	"regexp"
	"strings"

	"github.com/mbeaver502/rosalind_go/rosalind_lib/nucleotide"
)

// Rna represents a strand of RNA.
type Rna struct {
	Rna string
	rna []nucleotide.Nucleotide
}

// ErrInvalidRnaInput indicates that the provided string
// represents an invalid sequence of characters and
// cannot constitute a valid RNA sequence.
var ErrInvalidRnaInput = errors.New("rna: invalid input string")

// New creates a new Rna instance and returns a pointer to it.
func New(s string) (*Rna, error) {
	if !isValidString(s) {
		return nil, ErrInvalidRnaInput
	}

	su := strings.ToUpper(s)

	return &Rna{
		Rna: su,
		rna: toNucleotideSlice(su),
	}, nil
}

// NewFromDna creates a new Rna instance from the given Dna instance.
func TranscribeFromDna(d []nucleotide.Nucleotide) (*Rna, error) {
	nt := make([]nucleotide.Nucleotide, len(d))
	for i, n := range d {
		switch n {
		case nucleotide.Adenine, nucleotide.Cytosine, nucleotide.Guanine:
			nt[i] = n
		case nucleotide.Thymine:
			nt[i] = nucleotide.Uracil
		}
	}

	return &Rna{
		Rna: string(nt),
		rna: nt,
	}, nil
}

// Determine whether the given string contains only
// valid character representations of RNA nucleotides.
func isValidString(s string) bool {
	rex := regexp.MustCompile(`^[aAcCgGuU]+$`)
	return rex.MatchString(s)
}

// Convert the given string s into a slice of Nucleotide.
func toNucleotideSlice(s string) []nucleotide.Nucleotide {
	nts := make([]nucleotide.Nucleotide, len(s))
	for i, n := range s {
		nt := nucleotide.Nucleotide(n)
		switch nt {
		case nucleotide.Adenine, nucleotide.Cytosine, nucleotide.Guanine, nucleotide.Uracil:
			nts[i] = nt
		}
	}
	return nts
}

// String returns a string version of the Rna instance.
func (r *Rna) String() string {
	return r.Rna
}
