package giftstacking

import (
	"testing"
)

func TestSample1(t *testing.T) {
	input := [][]int{{3, 4},
		{0, 1},
		{6, 4},
		{2, 2}}
	expected := 3
	result := giftStacking(input)

	if result == expected {
		t.Log("passed", result)
	} else {
		t.Error("expected", expected, " got ", result)
	}
}

func TestSample2(t *testing.T) {
	input := [][]int{
		{3, 2},
		{1, 7},
		{7, 2},
	}
	expected := 2
	result := giftStacking(input)

	if result == expected {
		t.Log("passed", result)
	} else {
		t.Error("expected", expected, " got ", result)
	}
}

func TestSample5(t *testing.T) {
	input := [][]int{
		{3, 1},
		{2, 4},
		{4, 1},
	}
	expected := 3
	result := giftStacking(input)

	if result == expected {
		t.Log("passed", result)
	} else {
		t.Error("expected", expected, " got ", result)
	}
}
