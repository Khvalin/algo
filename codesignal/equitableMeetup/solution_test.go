package equitableMeetup

import (
	"testing"
)

func TestSample1(t *testing.T) {
	input := [][]int{{1, 2, 1, 1, 1, 1},
		{5, 4, 7, 4, 5, 9},
		{3, 3, 3, 3, 3, 3, 3, 3, 3}}
	expected := [3]int{4, 1, 2}
	result := equitableMeetup(input)
	output := [3]int{}
	copy(output[:], result)
	if output == expected {
		t.Log("passed", output)
	} else {
		t.Error("expected", expected, " got ", result)
	}
}
