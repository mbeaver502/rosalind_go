package nucleotide

import (
	"reflect"
	"testing"
)

func Test_Nucleotide(t *testing.T) {
	nt := Adenine
	if nt != 'A' {
		t.Errorf("FAIL: Adenine is not 'A'. want 'A'; got '%s'", string(nt))
	}
	nt = Cytosine
	if nt != 'C' {
		t.Errorf("FAIL: Cytosine is not 'C'. want 'C'; got '%s'", string(nt))
	}
	nt = Guanine
	if nt != 'G' {
		t.Errorf("FAIL: Guanine is not 'G'. want 'G'; got '%s'", string(nt))
	}
	nt = Thymine
	if nt != 'T' {
		t.Errorf("FAIL: Thymine is not 'T'. want 'T'; got '%s'", string(nt))
	}
	nt = Uracil
	if nt != 'U' {
		t.Errorf("FAIL: Uracil is not 'U'. want 'U'; got '%s'", string(nt))
	}
}

func Test_Nucleotide_ToSlice(t *testing.T) {
	tests := []struct {
		testName string
		input    string
		want     NucleotideSequence
	}{
		{
			testName: "valid nucleotide bases",
			input:    "ACGTU",
			want:     NucleotideSequence{Adenine, Cytosine, Guanine, Thymine, Uracil},
		},
		{
			testName: "invalid nucleotide bases",
			input:    "XYZ",
			want:     NucleotideSequence{},
		},
		{
			testName: "mix of valid and invalid",
			input:    "ACGXYZTU",
			want:     NucleotideSequence{Adenine, Cytosine, Guanine, Thymine, Uracil},
		},
		{
			testName: "empty input",
			input:    "",
			want:     NucleotideSequence{},
		},
	}

	for _, test := range tests {
		got := ToSlice(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("FAIL %s: slice lengths do not match. want %+v; got %+v", test.testName, test.want, got)
		}
	}
}
