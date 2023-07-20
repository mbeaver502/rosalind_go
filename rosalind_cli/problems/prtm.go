package problems

import (
	"fmt"
	"os"

	"github.com/mbeaver502/rosalind_go/rosalind_lib/protein"
)

type problemPrtm struct{}

// Do processes the prtm problem.
// https://rosalind.info/problems/prtm/
func (pp problemPrtm) Do(inFile, outFile *os.File) error {
	input := ""
	if inFile == os.Stdin {
		fmt.Println("Input:")
	}
	input = readInput(inFile)

	p, err := protein.New(input)
	if err != nil {
		return err
	}

	return writeOutput(fmt.Sprintf("%.3f\n", p.Mass()), outFile)
}
