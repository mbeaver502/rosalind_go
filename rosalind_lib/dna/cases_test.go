package dna

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
	wantMap    map[Nucleotide]uint
}

var dnaNewTests = []dnaTest{
	{
		testName:      "valid input string",
		input:         "ACGT",
		want:          &Dna{Dna: "ACGT", dna: []Nucleotide{NucleotideA, NucleotideC, NucleotideG, NucleotideT}},
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
		wantMap: map[Nucleotide]uint{
			NucleotideA: 1,
			NucleotideC: 1,
			NucleotideG: 1,
			NucleotideT: 1,
		},
	},
	{
		testName:   "medium string",
		input:      "AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC",
		wantString: "20 12 17 21",
		wantMap: map[Nucleotide]uint{
			NucleotideA: 20,
			NucleotideC: 12,
			NucleotideG: 17,
			NucleotideT: 21,
		},
	},
	{
		testName:   "long string",
		input:      "GACAAGTAAACTTGAACAAAGGGCCTGTACAGACGGGCTGTCTATAACGCCTGAGTTTGGTAGCCACGTAGGAGCCCGGGGCGCCTTTTAGGAGCAGACTGGTGGCGATCATCCTTGGTCTTACGATGTCTCGGCCTTATGAGTTATGAAAGGCTCCAGCTACTCCGAACAATTCATCGTTAACGTTAGTAGAGAACTAGACGGCATGCCGAGATGCAGTCAATACCGCTCCGTTGAGAGAGTGGTAGAATTGGTAATATACCGGAAGCAGAGCATTTCTTATCCTTGCTGACGCGTAGTTGTTGAAATCCTTGATCGACACTTCCCAGCTCGCAGCAATTTCAATGCCACTCCATCCGGCAAAATATATGAACACCAGGGAAACATACCGTTTGAACTGTACATTGCAGGGGCCCAGCGCATCGCAAATAGTCTCCCGAAGGTGGAGTTTCATAGAATAGATTTTGTTTAACTCACACGTCTGTCCGTGAAGCACATACACTGGCTGTCCTCCGCTGAATTTCACAACGCACCGGATGCTAGCATTGCGCGTTCAGTTGTACGCCCTCGAGTAACGGGGAGTAGGCAGCCCTACTGCTAACAAGTGGGTCTATATGAATGGACAGCAACCAGGCTCAAATCTGCCCTTTAGGAGACTTTGGAACGTACTGCCTATTAGTGTCACGATTATCCGCGGGAAGTACTAATAGGGCTAGGCAACTCCCAGTTTGGCTCGCAACGAATGTCAAGAAATAGAGGACATCCGGAGCAACAAGCCCCCGGGGCCACGCTTCTCCAAGCCCTCACCACCAGCACGTCCTGGCGTGCTGTAGGTATAGGGTGTTTCGTTCG",
		wantString: "219 215 216 202",
		wantMap: map[Nucleotide]uint{
			NucleotideA: 219,
			NucleotideC: 215,
			NucleotideG: 216,
			NucleotideT: 202,
		},
	},
}
