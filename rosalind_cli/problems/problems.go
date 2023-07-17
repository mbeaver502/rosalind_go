package problems

var RosalindProblems = map[string]func(){
	"dna":  problemDna,
	"rna":  problemRna,
	"revc": problemRevc,
}

func Do(prob *string) {
	// Problem is a required argument,
	// so no need to do a nil-check. Hopefully.
	RosalindProblems[*prob]()
}
