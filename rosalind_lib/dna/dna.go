package dna

import (
	"errors"
	"regexp"
	"strings"

	"github.com/mbeaver502/rosalind_go/rosalind_lib/nucleotide"
	"github.com/mbeaver502/rosalind_go/rosalind_lib/rna"
)

// Dna represents a strand of DNA.
type Dna struct {
	Dna string
	dna nucleotide.Sequence
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
		dna: nucleotide.ToSequence(su),
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
func (d *Dna) CountNucleotides() nucleotide.Counts {
	return d.countNucleotides()
}

// CountNucleotidesString counts the number of nucleotides in the Dna instance
// and returns the result in a pretty-printed string.
func (d *Dna) CountNucleotidesString() string {
	return d.CountNucleotides().String()
}

// Count the number of nucleotides in the Dna instance.
// If the given instance is nil, all counts default to
// zero values.
func (d *Dna) countNucleotides() nucleotide.Counts {
	counts := make(nucleotide.Counts)
	if d == nil {
		return counts
	}

	for _, n := range d.Dna {
		m := nucleotide.Nucleotide(n)
		switch m {
		case nucleotide.Adenine, nucleotide.Cytosine, nucleotide.Guanine, nucleotide.Thymine:
			counts[m]++
		}
	}
	return counts
}

// Transcribe produces an Rna instance from the given Dna instance.
func (d *Dna) Transcribe() (*rna.Rna, error) {
	return rna.TranscribeFromDna(d.dna)
}

func (d *Dna) ReverseComplement() nucleotide.Sequence {
	return reverseComplement(d.dna,
		nucleotide.Mapping{
			nucleotide.Adenine:  nucleotide.Thymine,
			nucleotide.Thymine:  nucleotide.Adenine,
			nucleotide.Cytosine: nucleotide.Guanine,
			nucleotide.Guanine:  nucleotide.Cytosine,
		})
}

func reverseComplement(ns nucleotide.Sequence, mapping nucleotide.Mapping) nucleotide.Sequence {
	rc := make(nucleotide.Sequence, len(ns))
	pos := 0
	for i := len(ns) - 1; i >= 0; i-- {
		rc[pos] = mapping.Transform(ns[i])
		pos++
	}
	return rc
}

// Iter returns an iterator on the internal Nucleotide Sequence.
func (d *Dna) Iter() *nucleotide.SequenceIterator {
	return d.dna.Iter()
}
