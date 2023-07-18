package problems

import (
	"fmt"
	"os"

	"github.com/mbeaver502/rosalind_go/rosalind_lib/dna"
)

type problemRevc struct{}

// https://rosalind.info/problems/revc/
func (pr problemRevc) Do(inFile, outFile *os.File) error {
	fmt.Println("Complementing a Strand of DNA")

	input := ""
	if inFile == os.Stdin {
		fmt.Println("Input:")
	}
	input = readInput(inFile)

	d, err := dna.New(input)
	if err != nil {
		return err
	}

	if outFile == os.Stdout {
		fmt.Println("Output:")
	}

	return writeOutput(d.ReverseComplement().String(), outFile)
}
