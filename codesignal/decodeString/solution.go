package main

import "fmt"

func decodeString(s string) string {
	type state struct {
		count int
		start int
	}

	res := []rune{}
	stack := []state{}
	c := 0

	for i := 0; i < len(s); i++ {
		ch := rune(s[i])
		switch {
		case ch >= '0' && ch <= '9':
			c *= 10
			c += int(ch - '0')
			break

		case ch == '[':
			stack = append(stack, state{c, i})
			c = 0
			break

		case ch == ']':
			l := len(stack) - 1
			stack[l].count--

			if stack[l].count > 0 {
				i = stack[l].start
			} else {
				stack = stack[:l]
			}
			break

		default:
			res = append(res, ch)
			break
		}
	}

	return string(res)
}

func main() {
	fmt.Println(decodeString("c2[a2[b]]"))
	fmt.Println(decodeString("c1[a1[b]]"))
	fmt.Println(decodeString("2[2[2[2[b]]]]"))
}
