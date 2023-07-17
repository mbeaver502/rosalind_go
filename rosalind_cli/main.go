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

func main() {
	parser := argparse.NewParser(progName, progDesc)

	problem := parser.String("p", "problem", &argparse.Options{
		Required: true,
		Validate: func(args []string) error {
			for _, arg := range args {
				if _, found := rosalindProblems[arg]; !found {
					return errors.New("invalid problem")
				}
			}
			return nil
		},
		Help:    `Specify the Rosalind problem in lowercase (identified at the end of the URL, e.g., https://rosalind.info/problems/dna/ => "dna")`,
		Default: nil,
	})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(parser.Usage(err))
		return
	}

	rosalindProblems[*problem]()
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
