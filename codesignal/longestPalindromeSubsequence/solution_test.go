package longestPalindromeSubsequence

import "testing"

func Test_Sample0(t *testing.T) {
	input := "ABCDAD"
	expected := "ABA"
	res := longestPalindromeSubsequence(input)
	if res != expected {
		t.Errorf("Test failed for input %s , expected %s, got %s ", input, expected, res)
	}
}

func Test_Sample1(t *testing.T) {
	input := "ABCCBDDCE"
	expected := "BCCB"
	res := longestPalindromeSubsequence(input)
	if res != expected {
		t.Errorf("Test failed for input %s , expected %s, got %s ", input, expected, res)
	}
}

func Test_Sample10(t *testing.T) {
	input := "ONLY ELEVEN CHARACTERS REMAIN"
	expected := "NEECARACEEN"
	res := longestPalindromeSubsequence(input)
	if res != expected {
		t.Errorf("Test failed for input %s , expected %s, got %s ", input, expected, res)
	}
}
