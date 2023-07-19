package problems

import (
	"fmt"
	"os"

	"github.com/mbeaver502/rosalind_go/rosalind_lib/dna"
)

type problemRevc struct{}

// Do processes the revc problem.
// https://rosalind.info/problems/revc/
func (pr problemRevc) Do(inFile, outFile *os.File) error {
	input := ""
	if inFile == os.Stdin {
		fmt.Println("Input:")
	}
	input = readInput(inFile)

	d, err := dna.New(input)
	if err != nil {
		return err
	}

	return writeOutput(d.ReverseComplement().String(), outFile)
}
