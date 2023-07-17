package dna

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	nt "github.com/mbeaver502/rosalind_go/rosalind_lib/nucleotide"
	"github.com/mbeaver502/rosalind_go/rosalind_lib/rna"
)

// Dna represents a strand of DNA.
type Dna struct {
	Dna string
	dna nt.NucleotideSequence
}

// NucleotideCounts maps the number of appearances of Nucleotides.
// Example: counts['A'] = 5
type NucleotideCounts map[nt.Nucleotide]uint

// NucleotideMapping maps one Nucleotide to another.
// Example: 'A' -> 'T'
type NucleotideMapping map[nt.Nucleotide]nt.Nucleotide

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
		dna: nt.ToSlice(su),
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
func (d *Dna) CountNucleotides() NucleotideCounts {
	return d.countNucleotides()
}

// CountNucleotidesString counts the number of nucleotides in the Dna instance
// and returns the result in a pretty-printed string.
func (d *Dna) CountNucleotidesString() string {
	counts := d.countNucleotides()
	return fmt.Sprintf(
		"%d %d %d %d",
		counts[nt.Adenine],
		counts[nt.Cytosine],
		counts[nt.Guanine],
		counts[nt.Thymine],
	)
}

// Count the number of nucleotides in the Dna instance.
// If the given instance is nil, all counts default to
// zero values.
func (d *Dna) countNucleotides() NucleotideCounts {
	counts := make(NucleotideCounts)
	if d == nil {
		return counts
	}

	for _, n := range d.Dna {
		m := nt.Nucleotide(n)
		switch m {
		case nt.Adenine, nt.Cytosine, nt.Guanine, nt.Thymine:
			counts[m]++
		}
	}
	return counts
}

// Transcribe produces an Rna instance from the given Dna instance.
func (d *Dna) Transcribe() (*rna.Rna, error) {
	return rna.TranscribeFromDna(d.dna)
}

func (d *Dna) ReverseComplement() nt.NucleotideSequence {
	return reverseComplement(d.dna,
		NucleotideMapping{
			nt.Adenine:  nt.Thymine,
			nt.Thymine:  nt.Adenine,
			nt.Cytosine: nt.Guanine,
			nt.Guanine:  nt.Cytosine,
		})
}

func reverseComplement(ns nt.NucleotideSequence, mapping NucleotideMapping) nt.NucleotideSequence {
	rc := make(nt.NucleotideSequence, len(ns))
	pos := 0
	for i := len(ns) - 1; i >= 0; i-- {
		rc[pos] = mapping[ns[i]]
		pos++
	}
	return rc
}
