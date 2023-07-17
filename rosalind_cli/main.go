package main

import (
	"log"

	"github.com/mbeaver502/rosalind_go/rosalind_lib/dna"
)

func main() {
	s := "ACGT"
	d := dna.New(s)

	log.Printf("input:  %s\n", d)
	log.Printf("output: %s\n", d.CountNucleotidesString())
}
