package luckykids

import "testing"

func Test_Sample1(t *testing.T) {
	input := []int{3, 3, 4, 5, 2, 1, 5, 5}
	expected := 2
	got := luckyKids(input)

	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func Test_Sample3(t *testing.T) {
	input := []int{10, 100, 1}
	expected := 3
	got := luckyKids(input)

	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
