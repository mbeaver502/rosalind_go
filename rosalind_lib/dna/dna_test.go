package dna

import (
	"testing"
)

func Test_Dna_New(t *testing.T) {
	for _, dt := range dnaNewTests {
		gotDna, gotErr := New(dt.input)
		if dt.errorExpected && gotErr == nil {
			t.Errorf("FAIL %s: expected error, but got none", dt.testName)
		}
		if gotErr == nil && !dt.errorExpected {
			if gotDna == nil {
				t.Errorf("FAIL %s: Dna instance came back nil", dt.testName)
			} else {
				if dt.want.Dna != gotDna.Dna {
					t.Errorf("FAIL %s: Dna.Dna does not match. want %s; got %s", dt.testName, dt.want.Dna, gotDna.Dna)
				}
				if gotDna.dna == nil {
					t.Errorf("FAIL %s: Dna.dna came back nil", dt.testName)
				} else {
					gotLen := len(gotDna.dna)
					wantLen := len(dt.want.dna)
					if wantLen != gotLen {
						t.Errorf("FAIL %s: dna len mismatch. want %d; got %d", dt.testName, wantLen, gotLen)
					} else {
						for i := 0; i < wantLen; i++ {
							if dt.want.dna[i] != gotDna.dna[i] {
								t.Errorf("FAIL %s: Dna.dna is incorrect. want %+v; got %+v", dt.testName, dt.want.dna, gotDna.dna)
							}
						}
					}
				}
			}
		}
	}
}

func Test_Dna_String(t *testing.T) {
	for _, dt := range dnaStringTests {
		d, err := New(dt.input)
		if err != nil {
			t.Errorf("FAIL %s: got error, but expected none", dt.testName)
		}

		if d.String() != dt.input {
			t.Errorf("FAIL %s: Dna string representation does not match input string. want %s; got %s", dt.testName, dt.want, dt.input)
		}
	}
}

func Test_Dna_CountNucleotides(t *testing.T) {
	for _, dt := range dnaCountNucleotidesTests {
		d, err := New(dt.input)
		if err != nil {
			t.Errorf("FAIL %s: got error, but expected none", dt.testName)
		}

		gotCNS := d.CountNucleotidesString()
		if dt.wantString != gotCNS {
			t.Errorf("FAIL %s: count nucleotides string does not match. want %s; got %s", dt.testName, dt.wantString, gotCNS)
		}

		gotCNM := d.CountNucleotides()
		if len(gotCNM) != len(dt.wantMap) {
			t.Errorf("FAIL %s: count nucleotides map does not match in length. want %+v; got %+v", dt.testName, dt.wantMap, gotCNM)
		}
		for k := range gotCNM {
			if gotCNM[k] != dt.wantMap[k] {
				t.Errorf("FAIL %s: count nucleotides map does not match on key %+v. want %d; got %d", dt.testName, k, dt.wantMap[k], gotCNM[k])
			}
		}
	}
}

func Test_Nucleotide(t *testing.T) {
	nt := NucleotideA
	if nt != 'A' {
		t.Errorf("FAIL: NucleotideA is not 'A'. want 'A'; got '%s'", string(nt))
	}
	nt = NucleotideC
	if nt != 'C' {
		t.Errorf("FAIL: NucleotideC is not 'C'. want 'C'; got '%s'", string(nt))
	}
	nt = NucleotideG
	if nt != 'G' {
		t.Errorf("FAIL: NucleotideG is not 'G'. want 'G'; got '%s'", string(nt))
	}
	nt = NucleotideT
	if nt != 'T' {
		t.Errorf("FAIL: NucleotideT is not 'T'. want 'T'; got '%s'", string(nt))
	}
}
