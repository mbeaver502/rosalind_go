package main

import (
	"fmt"
	"os"

	"github.com/mbeaver502/rosalind_go/rosalind_cli/problems"
)

func main() {
	parser := getParser()
	problem := addStringArgument(parser, "problem")
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(parser.Usage(err))
		return
	}

	problems.Do(problem)
}
