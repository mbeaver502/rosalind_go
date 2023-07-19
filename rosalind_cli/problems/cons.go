package problems

import (
	"errors"
	"fmt"
	"os"

	"github.com/mbeaver502/rosalind_go/rosalind_lib/utils/fasta"
	"github.com/mbeaver502/rosalind_go/rosalind_lib/utils/matrix"
)

type problemCons struct{}

// Do processes the cons problem.
// https://rosalind.info/problems/cons/
func (pc problemCons) Do(inFile, outFile *os.File) error {
	if inFile == os.Stdin {
		fmt.Println("Input:")
	}
	
	input := fasta.ReadFastaContent(inFile)
	if len(input) < 1 {
		return errors.New("gc: insufficient input")
	}

	ds, err := fasta.FastaToDna(input)
	if err != nil {
		return err
	}

	profile := matrix.CreateProfile(ds)
	return writeOutput(
		fmt.Sprintf("%s\n%s", profile.Consensus(), profile),
		outFile,
	)
}
