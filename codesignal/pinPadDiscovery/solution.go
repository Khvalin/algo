package main

import "fmt"

func pinPadDiscovery(logins [][]string) string {
	result := make([]rune, len(logins[0]))
	variants := make([]map[byte]bool, len(logins[0]))
	for i, seq := range logins {
		for j, pair := range seq {
			first, second := pair[0]-'0', pair[1]-'0'
			if second < first {
				first, second = second, first
			}
			if i == 0 {
				variants[j] = map[byte]bool{}
				variants[j][first], variants[j][second] = true, true
			}

			count := 0
			var d byte

			for k, v := range variants[j] {
				variants[j][k] = v && (k == first || k == second)
				if variants[j][k] {
					count++
					d = second
					if k == first {
						d = first
					}
				}
			}

			if i == len(logins)-1 {
				if count == 1 {
					result[j] = rune(d + '0')
				} else {
					result[j] = '?'
				}
			}
		}
	}

	return string(result)
}

func main() {
	logins := [][]string{{"23", "17", "58", "17"},
		{"39", "14", "05", "14"},
		{"37", "01", "59", "01"},
		{"37", "18", "56", "18"}}

	fmt.Println(pinPadDiscovery(logins))
}
