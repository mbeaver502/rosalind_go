package problems

import (
	"fmt"
	"os"

	"github.com/mbeaver502/rosalind_go/rosalind_lib/dna"
)

type problemDna struct{}

// Do processes the dna problem.
// https://rosalind.info/problems/dna/
func (pd problemDna) Do(inFile, outFile *os.File) error {
	fmt.Println("Counting DNA Nucleotides")

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

	return writeOutput(d.CountNucleotidesString(), outFile)
}
