package problems

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/mbeaver502/rosalind_go/rosalind_lib/utils"
)

type problemSubs struct{}

// Do processes the subs problem.
// https://rosalind.info/problems/subs/
func (ps problemSubs) Do(inFile, outFile *os.File) error {
	if inFile == os.Stdin {
		fmt.Println("Input:")
	}
	input := readLines(inFile, 2)

	if len(input) < 2 {
		return errors.New("subs: insufficient input")
	}

	indices := utils.FindAllSubstrings(input[0], input[1])
	var sb strings.Builder
	for _, idx := range indices {
		sb.WriteString(fmt.Sprintf("%d ", idx+1))
	}

	return writeOutput(fmt.Sprintf("%s\n", sb.String()), outFile)
}
