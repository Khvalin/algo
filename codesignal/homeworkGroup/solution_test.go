package main

import (
	"testing"
)

func TestSample1(t *testing.T) {
	t.Skip()
	input := [][]int{{-1, 2, 4, 3},
		{2, -1, 3, 4},
		{4, 3, -1, 3},
		{3, 4, 3, -1}}
	res := homeworkGroup(input)
	if res != 3 {
		t.Error(
			//	"For", input,
			"expected", 3,
			"got", res,
		)
	} else {
		t.Log("passed")
	}
}

func TestSample5(t *testing.T) {
	input := [][]int{
		{-1, 9, 29, 25},
		{9, -1, 8, 10},
		{29, 8, -1, 26},
		{25, 10, 26, -1}}

	res := homeworkGroup(input)
	if res != 25 {
		t.Error(
			//	"For", input,
			"expected", 25,
			"got", res,
		)
	}
}

/*
func TestSample19(t *testing.T) {
	input := [][]int{
		{-1, 106, 26, 26, 2, 13, 2, 53, 68, 2, 50, 17, 82, 13, 5, 18, 61, 26, 50, 2, 37, 25},
		{106, -1, 64, 32, 100, 53, 136, 13, 58, 80, 104, 41, 36, 85, 109, 40, 17, 100, 68, 80, 65, 81},
		{26, 64, -1, 32, 36, 5, 40, 45, 10, 16, 8, 25, 100, 53, 13, 8, 17, 4, 4, 16, 1, 1},
		{26, 32, 32, -1, 20, 13, 40, 5, 58, 16, 72, 1, 20, 13, 37, 8, 25, 52, 52, 16, 41, 41},
		{2, 100, 36, 20, -1, 17, 4, 45, 82, 4, 68, 13, 64, 5, 13, 20, 65, 40, 64, 4, 49, 37},
		{13, 53, 5, 13, 17, -1, 25, 26, 25, 5, 25, 8, 65, 26, 10, 1, 18, 13, 17, 5, 10, 8},
		{2, 136, 40, 40, 4, 25, -1, 73, 90, 8, 64, 29, 100, 17, 9, 32, 85, 36, 68, 8, 53, 37},
		{53, 13, 45, 5, 45, 26, 73, -1, 61, 37, 89, 10, 13, 32, 64, 17, 20, 73, 61, 37, 52, 58},
		{68, 58, 10, 58, 82, 25, 90, 61, -1, 50, 10, 53, 130, 101, 45, 26, 13, 18, 2, 50, 5, 13},
		{2, 80, 16, 16, 4, 5, 8, 37, 50, -1, 40, 9, 68, 13, 5, 8, 41, 20, 36, 0, 25, 17},
		{50, 104, 8, 72, 68, 25, 64, 89, 10, 40, -1, 61, 164, 97, 25, 32, 37, 4, 4, 40, 5, 5},
		{17, 41, 25, 1, 13, 8, 29, 10, 53, 9, 61, -1, 29, 10, 26, 5, 26, 41, 45, 9, 34, 32},
		{82, 36, 100, 20, 64, 65, 100, 13, 130, 68, 164, 29, -1, 37, 109, 52, 65, 136, 128, 68, 113, 117},
		{13, 85, 53, 13, 5, 26, 17, 32, 101, 13, 97, 10, 37, -1, 32, 25, 68, 65, 85, 13, 68, 58},
		{5, 109, 13, 37, 13, 10, 9, 64, 45, 5, 25, 26, 109, 32, -1, 17, 52, 9, 29, 5, 20, 10},
		{18, 40, 8, 8, 20, 1, 32, 17, 26, 8, 32, 5, 52, 25, 17, -1, 13, 20, 20, 8, 13, 13},
		{61, 17, 17, 25, 65, 18, 85, 20, 13, 41, 37, 26, 65, 68, 52, 13, -1, 37, 17, 41, 16, 26},
		{26, 100, 4, 52, 40, 13, 36, 73, 18, 20, 4, 41, 136, 65, 9, 20, 37, -1, 8, 20, 5, 1},
		{50, 68, 4, 52, 64, 17, 68, 61, 2, 36, 4, 45, 128, 85, 29, 20, 17, 8, -1, 36, 1, 5},
		{2, 80, 16, 16, 4, 5, 8, 37, 50, 0, 40, 9, 68, 13, 5, 8, 41, 20, 36, -1, 25, 17},
		{37, 65, 1, 41, 49, 10, 53, 52, 5, 25, 5, 34, 113, 68, 20, 13, 16, 5, 1, 25, -1, 2},
		{25, 81, 1, 41, 37, 8, 37, 58, 13, 17, 5, 32, 117, 58, 10, 13, 26, 1, 5, 17, 2, -1}}

	res := homeworkGroup(input)
	if res != 17 {
		t.Error(
			//	"For", input,
			"expected", 17,
			"got", res,
		)
	}
}

func TestSample4(t *testing.T) {
	input := [][]int{{-1, 188, 349, 61, 48, 4},
		{188, -1, 237, 50, 120, 155},
		{349, 237, -1, 262, 402, 341},
		{61, 50, 262, -1, 29, 40},
		{48, 120, 402, 29, -1, 29},
		{4, 155, 341, 40, 29, -1}}
	res := homeworkGroup(input)
	if res != 237 {
		t.Error(
			//	"For", input,
			"expected", 237,
			"got", res,
		)
	}
}*/