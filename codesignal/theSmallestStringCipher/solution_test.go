package theSmallestStringCipher

import "testing"

func TestSample_5(t *testing.T) {
	a := "b"
	b := "baa"
	ans := "baab"

	res := theSmallestStringCipher(a, b)
	if res != ans {
		t.Error( //	"For", input,
			"expected", ans,
			"got", res)
	}
}

func TestSample_Custom1(t *testing.T) {
	a := "cbccbcccba"
	b := "cbccbcccbc"
	ans := "cbcbccbccbcccbacccbc"

	res := theSmallestStringCipher(a, b)
	if res != ans {
		t.Error( //	"For", input,
			"expected", ans,
			"got", res)
	}
}
