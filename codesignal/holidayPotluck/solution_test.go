package holidayPotluck

import (
	"testing"
)

func TestSample1(t *testing.T) {
	input := [][]int{
		{0, 8},
		{4, 4},
		{10, 8},
		{8, 3},
		{5, 13},
	}
	capacity := 10
	expected := 96
	result := holidayPotluck(input, capacity)

	if expected == result {
		t.Log("passed", result)
	} else {
		t.Error("expected", expected, " got ", result)
	}
}
