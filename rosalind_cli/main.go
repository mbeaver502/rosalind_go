package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/akamensky/argparse"
)

const (
	progName = "rosalind"
	progDesc = "(Poor) Solutions to Rosalind exercises, written in Go."
)

var rosalindProblems = map[string]func(){
	"dna":  ProblemDna,
	"rna":  ProblemRna,
	"revc": ProblemRevc,
}

type argument struct {
	short   string
	long    string
	options *argparse.Options
}

var arguments = map[string]*argument{
	"problem": {
		short: "p",
		long:  "problem",
		options: &argparse.Options{
			Required: true,
			Validate: problemArgValidator,
			Help:     `Specify the Rosalind problem in lowercase (identified at the end of the URL, e.g., https://rosalind.info/problems/dna/ => "dna")`,
			Default:  nil,
		},
	},
}

func main() {
	parser := argparse.NewParser(progName, progDesc)

	probArg := arguments["problem"]
	problem := parser.String(probArg.short, probArg.long, probArg.options)

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(parser.Usage(err))
		return
	}

	// Problem is a required argument,
	// so need to do a nil-check. Hopefully.
	rosalindProblems[*problem]()
}

// Validate the provided problem argument
// by ensure it exists in our problem map.
func problemArgValidator(args []string) error {
	for _, arg := range args {
		if _, found := rosalindProblems[arg]; !found {
			return errors.New("invalid problem")
		}
	}
	return nil
}

// TODO: Implement DNA problem
// https://rosalind.info/problems/dna/
func ProblemDna() {
	log.Println("DNA problem")
}

// TODO: Implement RNA problem
// https://rosalind.info/problems/rna/
func ProblemRna() {
	log.Println("RNA problem")
}

// TODO: Implement Revc problem
// https://rosalind.info/problems/revc/
func ProblemRevc() {
	log.Println("Revc Problem")
}
