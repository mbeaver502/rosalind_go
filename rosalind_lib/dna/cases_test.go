package dna

import "github.com/mbeaver502/rosalind_go/rosalind_lib/nucleotide"

type dnaTest struct {
	testName      string
	input         string
	want          *Dna
	errorExpected bool
}

type dnaStringTest struct {
	testName string
	input    string
	want     string
}

type dnaCountNucleotidesTest struct {
	testName   string
	input      string
	wantString string
	wantMap    map[nucleotide.Nucleotide]uint
}

var dnaNewTests = []dnaTest{
	{
		testName: "valid input string",
		input:    "ACGT",
		want: &Dna{
			Dna: "ACGT",
			dna: []nucleotide.Nucleotide{
				nucleotide.Adenine,
				nucleotide.Cytosine,
				nucleotide.Guanine,
				nucleotide.Thymine,
			},
		},
		errorExpected: false,
	},
	{
		testName:      "invalid input string",
		input:         "ACGTb",
		want:          nil,
		errorExpected: true,
	},
	{
		testName:      "empty input string",
		input:         "",
		want:          nil,
		errorExpected: true,
	},
}

var dnaStringTests = []dnaStringTest{
	{
		testName: "string result matches input",
		input:    "ACGT",
		want:     "ACGT",
	},
}

var dnaCountNucleotidesTests = []dnaCountNucleotidesTest{
	{
		testName:   "simple string",
		input:      "ACGT",
		wantString: "1 1 1 1",
		wantMap: map[nucleotide.Nucleotide]uint{
			nucleotide.Adenine:  1,
			nucleotide.Cytosine: 1,
			nucleotide.Guanine:  1,
			nucleotide.Thymine:  1,
		},
	},
	{
		testName:   "medium string",
		input:      "AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC",
		wantString: "20 12 17 21",
		wantMap: map[nucleotide.Nucleotide]uint{
			nucleotide.Adenine:  20,
			nucleotide.Cytosine: 12,
			nucleotide.Guanine:  17,
			nucleotide.Thymine:  21,
		},
	},
	{
		testName:   "long string",
		input:      "GACAAGTAAACTTGAACAAAGGGCCTGTACAGACGGGCTGTCTATAACGCCTGAGTTTGGTAGCCACGTAGGAGCCCGGGGCGCCTTTTAGGAGCAGACTGGTGGCGATCATCCTTGGTCTTACGATGTCTCGGCCTTATGAGTTATGAAAGGCTCCAGCTACTCCGAACAATTCATCGTTAACGTTAGTAGAGAACTAGACGGCATGCCGAGATGCAGTCAATACCGCTCCGTTGAGAGAGTGGTAGAATTGGTAATATACCGGAAGCAGAGCATTTCTTATCCTTGCTGACGCGTAGTTGTTGAAATCCTTGATCGACACTTCCCAGCTCGCAGCAATTTCAATGCCACTCCATCCGGCAAAATATATGAACACCAGGGAAACATACCGTTTGAACTGTACATTGCAGGGGCCCAGCGCATCGCAAATAGTCTCCCGAAGGTGGAGTTTCATAGAATAGATTTTGTTTAACTCACACGTCTGTCCGTGAAGCACATACACTGGCTGTCCTCCGCTGAATTTCACAACGCACCGGATGCTAGCATTGCGCGTTCAGTTGTACGCCCTCGAGTAACGGGGAGTAGGCAGCCCTACTGCTAACAAGTGGGTCTATATGAATGGACAGCAACCAGGCTCAAATCTGCCCTTTAGGAGACTTTGGAACGTACTGCCTATTAGTGTCACGATTATCCGCGGGAAGTACTAATAGGGCTAGGCAACTCCCAGTTTGGCTCGCAACGAATGTCAAGAAATAGAGGACATCCGGAGCAACAAGCCCCCGGGGCCACGCTTCTCCAAGCCCTCACCACCAGCACGTCCTGGCGTGCTGTAGGTATAGGGTGTTTCGTTCG",
		wantString: "219 215 216 202",
		wantMap: map[nucleotide.Nucleotide]uint{
			nucleotide.Adenine:  219,
			nucleotide.Cytosine: 215,
			nucleotide.Guanine:  216,
			nucleotide.Thymine:  202,
		},
	},
}
