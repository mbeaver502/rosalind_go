package nucleotide

import "fmt"

// Counts maps the number of appearances of Nucleotides.
// Example: counts['A'] = 5
type Counts map[Nucleotide]uint

// String returns a pretty printed representation of the NucleotideCounts.
// Currently only returns DNA nucleotides.
func (nc Counts) String() string {
	return fmt.Sprintf(
		"%d %d %d %d\n",
		nc[Adenine],
		nc[Cytosine],
		nc[Guanine],
		nc[Thymine],
	)
}
