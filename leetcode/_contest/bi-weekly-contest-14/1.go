package main

import (
	"fmt"
	"strconv"
	"strings"
)

func toHexspeak(num string) string {
	n, _ := strconv.Atoi(num)
	s := strings.ToUpper(fmt.Sprintf("%0x", n))
	res := ""

	for _, ch := range s {
		if ch == '1' {
			res += "I"
			continue
		}

		if ch == '0' {
			res += "O"
			continue
		}

		if ch < 'A' {
			return "ERROR"
		}

		res += string(ch)
	}
	return res
}

func main() {
	fmt.Println(toHexspeak("257"))
	fmt.Println(toHexspeak("747823223228"))
}
