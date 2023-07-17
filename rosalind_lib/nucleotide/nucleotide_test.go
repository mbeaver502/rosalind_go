package nucleotide

import "testing"

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
