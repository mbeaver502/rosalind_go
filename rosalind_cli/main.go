package main

import (
	"log"

	"github.com/mbeaver502/rosalind_go/rosalind_lib/dna"
)

func main() {
	s := "ACGTaCgTAcGt"
	d, err := dna.New(s)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("input:  %s\n", d)
	log.Printf("output: %s\n", d.CountNucleotidesString())
}
