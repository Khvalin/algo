package wordLadder

import (
	"testing"
)

func TestSampleCustom1(t *testing.T) {
	t.Skip()
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	res := ladderLength(beginWord, endWord, wordList)
	if 5 == res {
		t.Log("Test 1 passed")
	} else {
		t.Error("fail, got ", res, " expected ", 5)
	}
}

func TestSampleCustom2(t *testing.T) {
	t.Skip()
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log"}
	res := ladderLength(beginWord, endWord, wordList)
	if 0 == res {
		t.Log("Test 1 passed")
	} else {
		t.Error("fail, got ", res, " expected ", 0)
	}
}

func TestSampleCustom30(t *testing.T) {
	//	t.Skip()
	beginWord := "qa"
	endWord := "sq"
	wordList := []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"}
	res := ladderLength(beginWord, endWord, wordList)
	if 5 == res {
		t.Log("Test 1 passed")
	} else {
		t.Error("fail, got ", res, " expected ", 0)
	}
}
