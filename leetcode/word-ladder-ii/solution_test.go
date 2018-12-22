package wordLadder

import (
	"reflect"
	"testing"
)

func TestSampleCustom1(t *testing.T) {
	//	t.Skip()
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	res := findLadders(beginWord, endWord, wordList)
	answer := [][]string{{"hit", "hot", "dot", "dog", "cog"},
		{"hit", "hot", "lot", "log", "cog"}}

	if reflect.DeepEqual(res, answer) {
		t.Log("Test 1 passed")
	} else {
		t.Error("fail, got ", res, " expected ", answer)
	}
}
