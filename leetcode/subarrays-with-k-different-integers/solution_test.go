package subarraysWithKDistinct

import (
	"testing"
)

func TestSampleCustom1(t *testing.T) {
	A, K := []int{1, 2, 1, 2, 3}, 2
	res := subarraysWithKDistinct(A, K)
	if res == 7 {
		t.Log("OK")
	} else {
		t.Error("fail, got ", res, " expected ", 7)
	}
}

func TestSampleCustom2(t *testing.T) {
	A, K := []int{1, 2, 1, 3, 4}, 3
	res := subarraysWithKDistinct(A, K)
	if res == 3 {
		t.Log("OK")
	} else {
		t.Error("fail, got ", res, " expected ", 3)
	}
}
