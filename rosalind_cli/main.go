package main

import (
	"log"

	"github.com/mbeaver502/rosalind_go/rosalind_lib/dna"
)

func main() {
	s := "AAAACCCGGT"
	d, err := dna.New(s)
	if err != nil {
		log.Fatalln(err)
	}

	r, err := d.Transcribe()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("dna: %s\n", d)
	log.Printf("rna: %s\n", r)
	log.Printf("rc:  %s\n", d.ReverseComplement())
}
