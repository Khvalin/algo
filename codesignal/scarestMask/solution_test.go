package main

import "testing"

func TestSample1(t *testing.T) {
	input := []string{"A    A",
		" O  O ",
		"= WW ="}

	res := scariestMask(input)
	if res != 0 {
		t.Error( //	"For", input,
			"expected", 0,
			"got", res)
	}
}
