package main

import (
	"errors"

	"github.com/akamensky/argparse"
	"github.com/mbeaver502/rosalind_go/rosalind_cli/problems"
)

const (
	progName = "rosalind"
	progDesc = "(Poor) Solutions to Rosalind exercises, written in Go."
)

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
	"inFile": {
		short: "i",
		long:  "input",
		options: &argparse.Options{
			Required: false,
			Validate: nil,
			Help:     `Specify an input file. Either omit or specify "-" to use stdin`,
			Default:  "-",
		},
	},
	"outFile": {
		short: "o",
		long:  "output",
		options: &argparse.Options{
			Required: false,
			Validate: nil,
			Help:     `Specify an output file. Either omit or specify "-" to use stdout`,
			Default:  "-",
		},
	},
}

// Validate the provided problem argument
// by ensuring it exists in our problem map.
func problemArgValidator(args []string) error {
	for _, arg := range args {
		if _, found := problems.RosalindProblems[arg]; !found {
			return errors.New("invalid problem")
		}
	}
	return nil
}

func getParser() *argparse.Parser {
	return argparse.NewParser(progName, progDesc)
}

func addStringArgument(p *argparse.Parser, arg string) *string {
	argObj := arguments[arg]
	return p.String(argObj.short, argObj.long, argObj.options)
}
