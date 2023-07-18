package problems

import (
	"fmt"
	"os"

	"github.com/mbeaver502/rosalind_go/rosalind_lib/dna"
)

type problemRna struct{}

// https://rosalind.info/problems/rna/
func (pr problemRna) Do(inFile, outFile *os.File) error {
	fmt.Println("Transcribing DNA into RNA")

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

	r, err := d.Transcribe()
	if err != nil {
		return err
	}

	return writeOutput(r.String(), outFile)
}
