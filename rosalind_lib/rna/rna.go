package rna

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/mbeaver502/rosalind_go/rosalind_lib/nucleotide"
)

// Rna represents a strand of RNA.
type Rna struct {
	Rna string
	rna nucleotide.Sequence
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
		rna: nucleotide.ToSequence(su),
	}, nil
}

// TranscribeFromDna creates a new Rna instance from the given slice of Nucleotides.
func TranscribeFromDna(d nucleotide.Sequence) (*Rna, error) {
	transcribed := make(nucleotide.Sequence, len(d))
	for i, n := range d {
		switch n {
		case nucleotide.Adenine, nucleotide.Cytosine, nucleotide.Guanine:
			transcribed[i] = n
		case nucleotide.Thymine:
			transcribed[i] = nucleotide.Uracil
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
	return fmt.Sprintf("%s\n", r.Rna)
}
