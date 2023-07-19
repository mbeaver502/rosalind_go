package matrix

import (
	"fmt"
	"strings"

	"github.com/mbeaver502/rosalind_go/rosalind_lib/dna"
	"github.com/mbeaver502/rosalind_go/rosalind_lib/nucleotide"
)

// Profile represents a profile matrix that
// maps Nucleotide to counts.
type Profile map[nucleotide.Nucleotide][]int

// NewProfile initializes a new Profile where
// each component slice is of the given size.
func NewProfile(size int) Profile {
	profile := make(Profile)
	profile[nucleotide.Adenine] = make([]int, size)
	profile[nucleotide.Cytosine] = make([]int, size)
	profile[nucleotide.Guanine] = make([]int, size)
	profile[nucleotide.Thymine] = make([]int, size)

	return profile
}

// CreateProfile creates a Profile from the given slice
// of pointers to dna.Dna instances.
func CreateProfile(ds []*dna.Dna) Profile {
	numRows := len(ds)
	numCols := len(ds[0].Dna)
	profile := NewProfile(numCols)

	for c := 0; c < numCols; c++ {
		for r := 0; r < numRows; r++ {
			nt := nucleotide.Nucleotide(ds[r].Dna[c])
			profile[nt][c]++
		}
	}

	return profile
}

// Consensus builds a consensus string from the Profile p.
func (p Profile) Consensus() string {
	var sb strings.Builder
	sA := p[nucleotide.Adenine]
	sC := p[nucleotide.Cytosine]
	sG := p[nucleotide.Guanine]
	sT := p[nucleotide.Thymine]

	for i := 0; i < len(sA); i++ {
		if sA[i] >= sC[i] && sA[i] >= sG[i] && sA[i] >= sT[i] {
			sb.WriteRune(rune(nucleotide.Adenine))
		} else if sC[i] >= sA[i] && sC[i] >= sG[i] && sC[i] >= sT[i] {
			sb.WriteRune(rune(nucleotide.Cytosine))
		} else if sG[i] >= sA[i] && sG[i] >= sC[i] && sG[i] >= sT[i] {
			sb.WriteRune(rune(nucleotide.Guanine))
		} else if sT[i] >= sA[i] && sT[i] >= sC[i] && sT[i] >= sG[i] {
			sb.WriteRune(rune(nucleotide.Thymine))
		}
	}

	return sb.String()
}

// String creates a pretty-printed representation
// of the Profile p.
func (p Profile) String() string {
	var sb strings.Builder
	for _, v := range p[nucleotide.Adenine] {
		sb.WriteString(fmt.Sprintf("%d ", v))
	}
	strA := strings.TrimSpace(sb.String())
	sb.Reset()

	for _, v := range p[nucleotide.Cytosine] {
		sb.WriteString(fmt.Sprintf("%d ", v))
	}
	strC := strings.TrimSpace(sb.String())
	sb.Reset()

	for _, v := range p[nucleotide.Guanine] {
		sb.WriteString(fmt.Sprintf("%d ", v))
	}
	strG := strings.TrimSpace(sb.String())
	sb.Reset()

	for _, v := range p[nucleotide.Thymine] {
		sb.WriteString(fmt.Sprintf("%d ", v))
	}
	strT := strings.TrimSpace(sb.String())
	sb.Reset()

	return fmt.Sprintf("%s: %s\n%s: %s\n%s: %s\n%s: %s\n",
		nucleotide.Adenine,
		strA,
		nucleotide.Cytosine,
		strC,
		nucleotide.Guanine,
		strG,
		nucleotide.Thymine,
		strT,
	)
}
