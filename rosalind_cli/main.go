package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mbeaver502/rosalind_go/rosalind_cli/problems"
)

func main() {
	parser := getParser()
	problem := addStringArgument(parser, "problem")
	inFile := addStringArgument(parser, "inFile")
	outFile := addStringArgument(parser, "outFile")

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(parser.Usage(err))
		return
	}

	err = problems.Do(problem, inFile, outFile)
	if err != nil {
		log.Fatalln(err)
	}
}
