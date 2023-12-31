package problems

import (
	"errors"
	"fmt"
	"os"

	"github.com/mbeaver502/rosalind_go/rosalind_lib/dna"
	"github.com/mbeaver502/rosalind_go/rosalind_lib/utils/fasta"
)

type problemGc struct{}

type maxFastaEntry struct {
	entry     fasta.FastaEntry
	gcContent float64
}

// Do processes the gc problem.
// https://rosalind.info/problems/gc/
func (pg problemGc) Do(inFile, outFile *os.File) error {
	if inFile == os.Stdin {
		fmt.Println("Input:")
	}
	input := fasta.ReadFastaContent(inFile)

	if len(input) < 1 {
		return errors.New("gc: insufficient input")
	}

	maxItem, err := pg.getMaxFasta(input)
	if err != nil {
		return err
	}

	return writeOutput(
		fmt.Sprintf("%s\n%f\n", maxItem.entry.Label, maxItem.gcContent),
		outFile,
	)
}

func (pg problemGc) getMaxFasta(input fasta.FastaContent) (*maxFastaEntry, error) {
	d, err := dna.New(input[0].Content)
	if err != nil {
		return nil, err
	}

	maxItem := &maxFastaEntry{
		entry:     input[0],
		gcContent: d.GcContent(),
	}

	for _, item := range input[1:] {
		d, err := dna.New(item.Content)
		if err != nil {
			return nil, err
		}
		thisGc := d.GcContent()
		if thisGc > maxItem.gcContent {
			maxItem.entry.Label = item.Label
			maxItem.gcContent = thisGc
		}
	}

	return maxItem, nil
}
