package main

import "fmt"

type trie struct {
	char   byte
	next   *trie
	isWord bool
}

type StreamChecker struct {
	tries  []*trie
	maxLen int
	q      []byte
}

func Constructor(words []string) StreamChecker {
	tries := make([]*trie, len(words))
	maxLen := 0

	for i, w := range words {
		if len(w) > maxLen {
			maxLen = len(w)
		}
		var prev *trie

		for j := len(w); j >= 0; j-- {
			t := &trie{}

			if prev == nil {
				tries[i] = t
			} else {
				ch := byte(w[j])
				prev.char = ch
				prev.next = t
			}

			prev = t
		}

		prev.isWord = true
	}

	return StreamChecker{tries, maxLen, nil}
}

func (this *StreamChecker) Query(letter byte) bool {
	this.q = append(this.q, letter)
	if len(this.q) > this.maxLen {
		this.q = this.q[1:]
	}

	candidates := []*trie{}
	for _, t := range this.tries {
		if t.char == letter && t.next != nil {
			candidates = append(candidates, t)
		}
	}

	for i := len(this.q) - 1; i >= 0; i-- {
		for j := 0; j < len(candidates); {
			c := candidates[j]
			if c != nil && c.char == this.q[i] && c.next != nil {
				if c.next.isWord {
					return true
				}

				candidates[j] = c.next
				j++
				continue
			}

			candidates[j] = candidates[len(candidates)-1]
			candidates = candidates[:len(candidates)-1]
		}
	}

	return false
}

/**
* Your StreamChecker object will be instantiated and called as such:
* obj := Constructor(words);
* param_1 := obj.Query(letter);
 */

func main() {
	obj := Constructor([]string{"ab", "ba", "aaab", "abab", "baa"})
	fmt.Println(obj.Query('a'))
	fmt.Println(obj.Query('b'))
}
