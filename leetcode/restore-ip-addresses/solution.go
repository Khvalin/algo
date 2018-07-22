package main

import "fmt"

func restoreIpAddresses(s string) []string {
	result := []string{}
	dotsIndices := []int{1, 2, 3, len(s)}

	splitByDots := func() []string {
		ans := []string{}
		prev := (0)
		for _, index := range dotsIndices {
			octet := ""
			if prev < index && index <= len(s) {
				octet =
					s[prev:index]
			}

			ans = append(ans, octet)
			prev = index
		}

		return ans
	}

	result = splitByDots()

	return result
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("0"))
}
