package sugoku

import (
	"testing"
)

func TestContains(t *testing.T) {

	empty := []int{}
	var val int
	val = 1
	// Try with empty list
	result := contains(empty, val)

	if result == true {
		t.Errorf("contains expected false for %v in %v got %v", val, empty, result)
	}

	// Try with list of values
	vC := []int{1, 2, 3}
	val = 3
	result = contains(vC, val)

	if result == false {
		t.Errorf("contains expected true for %v in %v got %v", val, vC, result)
	}

	val = 4
	result = contains(vC, val)
	if result == true {
		t.Errorf("contains expected true for %v in %v got %v", val, vC, result)
	}
}

func TestCleanGrid(t *testing.T) {
	var s Sugoku
	s[6][6] = 7

	cleanGrid(&s, 7)

	if s[6][6] != 0 {
		t.Errorf("cleanGrid expected 0 got %v", s[6][6])
	}

	s[0][0] = 2
	cleanGrid(&s, 1)

	if s[0][0] != 0 {
		t.Errorf("cleanGrid expected 0 got %v", s[0][0])
	}

	// Cleaning the wrong grid
	s[5][5] = 5
	cleanGrid(&s, 6)

	if s[5][5] == 0 {
		t.Errorf("cleanGrid expected 5 got %v", s[5][5])
	}
}

func TestGetPossibleValues(t *testing.T) {
	var s Sugoku
	var vG []int

	pV := getPossibleValues(0, 0, vG, &s)

	if len(pV) != limitNumber {
		t.Errorf("getPossibleValues expected %v values but got %v", limitNumber, len(pV))
	}

	for i := 1; i <= limitNumber; i++ {
		if !contains(pV, i) {
			t.Errorf("getPossibleValues expected %v in %v", i, pV)
		}
	}

	s[0][0] = 1
	s[1][2] = 3
	vG = append(vG, 5)
	pV = getPossibleValues(0, 3, vG, &s)

	if len(pV) != 7 {
		t.Errorf("getPossibleValues expected %v values but got %v", 7, len(pV))
	}

	if contains(pV, 1) && contains(pV, 5) {
		t.Errorf("getPossibleValues did not expected %v or %v but got %v", 1, 5, pV)
	}
}

func TestValidateCondition(t *testing.T) {
	var s Sugoku
	s[0][0] = 1
	s[1][2] = 5

	cond := validateCondition(&s, 0, 2, 1)

	if cond == false {
		t.Errorf("validateCondition should return true but returned %v", cond)
	}

	cond = validateCondition(&s, 0, 2, 6)
	if cond == true {
		t.Errorf("validateCondition should return false but returned %v", cond)
	}

	cond = validateCondition(&s, 0, 2, 5)
	if cond == false {
		t.Errorf("validateCondition should return true but returned %v", cond)
	}
}

func TestSolve(t *testing.T) {
	var s Sugoku
	s.Solve()
}

func TestGenerateSugoku(t *testing.T) {
	s := GenerateSugoku()

	for i := 0; i < limitNumber; i++ {
		if contains(s[i][:], 0) == true {
			t.Errorf("GenerateSugoku should not contain 0 in row #%v. Values %v", i, s[i])
		}
	}
}
