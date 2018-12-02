package medianScoresSolution

import (
	"testing"
)

func TestSample1(t *testing.T) {
	expected := [5]int{100, 60, 50, 60, 50}
	input := []int{100, 20, 50, 70, 45}
	result := medianScores(input)
	output := [5]int{}

	copy(output[:], result)

	if output == expected {
		t.Log("passed", output)
	} else {
		t.Error("failed! got ", output)
	}
}
