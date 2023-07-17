package nucleotide

import (
	"fmt"
	"strings"
)

// Mapping maps one Nucleotide to another.
// Example: 'A' -> 'T'
type Mapping map[Nucleotide]Nucleotide

// String returns a pretty printed representation of the NucleotideMapping.
func (nm Mapping) String() string {
	var sb strings.Builder
	for k, v := range nm {
		sb.WriteString(fmt.Sprintf("%s -> %s, ", k, v))
	}
	return fmt.Sprintf("{%s}", strings.TrimRight(sb.String(), ", "))
}
