package rna

import (
	"errors"
	"regexp"
	"strings"

	nt "github.com/mbeaver502/rosalind_go/rosalind_lib/nucleotide"
)

// Rna represents a strand of RNA.
type Rna struct {
	Rna string
	rna nt.NucleotideSequence
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
		rna: nt.ToSlice(su),
	}, nil
}

// NewFromDna creates a new Rna instance from the given slice of Nucleotides.
func TranscribeFromDna(d nt.NucleotideSequence) (*Rna, error) {
	transcribed := make(nt.NucleotideSequence, len(d))
	for i, n := range d {
		switch n {
		case nt.Adenine, nt.Cytosine, nt.Guanine:
			transcribed[i] = n
		case nt.Thymine:
			transcribed[i] = nt.Uracil
		}
	}

	return &Rna{
		Rna: string(transcribed),
		rna: transcribed,
	}, nil
}

// Determine whether the given string contains only
// valid character representations of RNA nucleotides.
func isValidString(s string) bool {
	rex := regexp.MustCompile(`^[aAcCgGuU]+$`)
	return rex.MatchString(s)
}

// String returns a string version of the Rna instance.
func (r *Rna) String() string {
	return r.Rna
}
