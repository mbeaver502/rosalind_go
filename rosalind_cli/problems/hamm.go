package problems

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mbeaver502/rosalind_go/rosalind_lib/utils"
)

type problemHamm struct{}

// Do processes the hamm problem.
// https://rosalind.info/problems/hamm/
func (ph problemHamm) Do(inFile, outFile *os.File) error {
	fmt.Println("Counting Point Mutations")

	if inFile == os.Stdin {
		fmt.Println("Input:")
	}
	input := readLines(inFile, 2)

	if outFile == os.Stdout {
		fmt.Println("Output:")
	}

	return writeOutput(
		fmt.Sprintf("%d\n", utils.HammingDistance(input[0], input[1])),
		outFile,
	)
}

func readLines(f *os.File, maxLines int) []string {
	lines := []string{}
	scanner := bufio.NewScanner(f)
	count := 0
	for count < maxLines && scanner.Scan() {
		lines = append(lines, scanner.Text())
		count++
	}
	return lines
}
