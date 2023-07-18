package problems

import (
	"bufio"
	"os"
)

type rosalindProblem interface {
	Do(inFile, outFile *os.File) error
}

var RosalindProblems = map[string]rosalindProblem{
	"dna":  problemDna{},
	"rna":  problemRna{},
	"revc": problemRevc{},
}

func Do(prob, inFile, outFile *string) error {
	input, err := getInFile(inFile)
	if err != nil {
		return err
	}
	defer input.Close()

	output, err := getOutFile(outFile)
	if err != nil {
		return err
	}
	defer output.Close()

	// Problem is a required argument,
	// so no need to do a nil-check. Hopefully.
	return RosalindProblems[*prob].Do(input, output)
}

func getInFile(f *string) (*os.File, error) {
	if f == nil || *f == "-" {
		return os.Stdin, nil
	}

	file, err := os.OpenFile(*f, os.O_RDONLY, os.ModeType)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func getOutFile(f *string) (*os.File, error) {
	if f == nil || *f == "-" {
		return os.Stdout, nil
	}

	file, err := os.Create(*f)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func readInput(f *os.File) string {
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	return scanner.Text()
}

func writeOutput(output string, f *os.File) error {
	w := bufio.NewWriter(f)
	if _, err := w.WriteString(output); err != nil {
		return err
	}
	return w.Flush()
}
