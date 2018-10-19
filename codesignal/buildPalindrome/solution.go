package main

import (
	"fmt"
)

func buildPalindrome(st string) string {
	p, i, j := -1, 0, len(st)-1
	for i < j {
		d := 1
		if st[i] != st[j] {
			p = i
			if j < len(st)-1 {
				j = len(st) - 1
				d = 0
			}
		} else {
			if p >= i {
				p = i - 1
			}
			j--
		}

		i += d
	}
	for ; p >= 0; p-- {
		st += string(st[p])
	}

	return st
}

func main() {
	fmt.Println(buildPalindrome("abcdc"), "\nabcdcba\n")
	fmt.Println("abccdc", buildPalindrome("abccdc"), "\nabccdccba")
}
