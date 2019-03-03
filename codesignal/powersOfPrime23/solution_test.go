package powersOfPrime23

import (
	"testing"
)

func TestSample1(t *testing.T) {
	expected := [2]int{6, 3}
	input := 37
	result := powersOfPrime23(input)
	output := [2]int{}

	copy(output[:], result)

	if output == expected {
		t.Log("passed", output)
	} else {
		t.Error("failed! got ", output)
	}
}

func TestSample2(t *testing.T) {
	input := 113
	result := powersOfPrime23(input)

	if len(result) == 0 {
		t.Log("passed", result)
	} else {
		t.Error("failed! got ", result)
	}
}

func TestSample_Custom1(t *testing.T) {
	expected := [2]int{0, 0}
	input := 0
	result := powersOfPrime23(input)
	output := [2]int{}

	copy(output[:], result)

	if output == expected {
		t.Log("passed", output)
	} else {
		t.Error("failed! got ", output)
	}
}

func TestSample_Custom2(t *testing.T) {
	expected := [2]int{1, 0}
	input := 1
	result := powersOfPrime23(input)
	output := [2]int{}

	copy(output[:], result)

	if output == expected {
		t.Log("passed", output)
	} else {
		t.Error("failed! got ", output)
	}
}

func TestSample_Custom3(t *testing.T) {
	expected := [2]int{0, 0}
	input := 0
	result := powersOfPrime23(input)
	output := [2]int{}

	copy(output[:], result)

	if output == expected {
		t.Log("passed", output)
	} else {
		t.Error("failed! got ", output)
	}
}
