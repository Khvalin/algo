package booAnalysis

import (
	"testing"
)

func TestSampleCustom1(t *testing.T) {
	input := "ooouuoo"
	res := booAnalysis(input)
	if res == 6 {
		t.Log("OK")
	} else {
		t.Error("fail, got ", res, " expected ", 6)
	}
}

func TestSampleCustom2(t *testing.T) {
	input := "ooouu"
	res := booAnalysis(input)
	if res == 0 {
		t.Log("OK")
	} else {
		t.Error("fail, got ", res, " expected ", 0)
	}
}

func TestSampleCustom6(t *testing.T) {
	input := "ouuooouo"
	res := booAnalysis(input)
	if res == 7 {
		t.Log("OK")
	} else {
		t.Error("fail, got ", res, " expected ", 7)
	}
}

func TestSampleCustom9(t *testing.T) {
	input := "ooouooouo"
	res := booAnalysis(input)
	if res == 8 {
		t.Log("OK")
	} else {
		t.Error("fail, got ", res, " expected ", 8)
	}
}
