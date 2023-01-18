package interview

import "testing"

func TestWordSearch(t *testing.T) {
	m := [][]uint8{
		{'A', 'U', 'I', 'K', 'F', 'W', 'N'},
		{'W', 'Q', 'B', 'O', 'L', 'X', 'P'},
		{'T', 'L', 'A', 'E', 'R', 'E', 'S'},
		{'Y', 'Z', 'X', 'E', 'R', 'L', 'W'},
	}

	s := "UBER"

	if !WordSearch(s, m) {
		t.Fatalf("For word %s where expected true output false", s)
	}

	s = "XLW"

	if !WordSearch(s, m) {
		t.Fatalf("For word %s where expected true output false", s)
	}

	s = "AERO"

	if WordSearch(s, m) {
		t.Fatalf("For word %s where expected false output true", s)
	}

}
