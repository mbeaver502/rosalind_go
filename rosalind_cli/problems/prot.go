package problems

import (
	"fmt"
	"os"

	"github.com/mbeaver502/rosalind_go/rosalind_lib/rna"
)

type problemProt struct{}

// Do processes the prot problem.
// https://rosalind.info/problems/prot/
func (pp problemProt) Do(inFile, outFile *os.File) error {
	input := ""
	if inFile == os.Stdin {
		fmt.Println("Input:")
	}
	input = readInput(inFile)

	r, err := rna.New(input)
	if err != nil {
		return err
	}

	return writeOutput(r.Translate().String(), outFile)
}
