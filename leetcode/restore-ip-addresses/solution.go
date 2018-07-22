package main

import (
	"fmt"
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	result := []string{}
	octetLengths := []byte{1, 1, 1, 1}

	for octetLengths[0] < 4 {
		var total int
		for _, val := range octetLengths {
			total += int(val)
		}
		if total == len(s) {
			//fmt.Println(octetLengths)

			var prev byte
			valid := true
			octets := []string{}

			for _, val := range octetLengths {
				fragment := s[prev : prev+val]

				octet, _ := strconv.Atoi(fragment)
				valid = valid && octet < 256 && len(fragment) > 0 && (len(fragment) == 1 || fragment[0] != '0')
				if !valid {
					break
				}
				octets = append(octets, fragment)

				prev += val
			}

			if valid {
				result = append(result, strings.Join(octets, "."))
			}
		}

		cur := len(octetLengths) - 1
		octetLengths[cur]++
		for cur > 0 && octetLengths[cur] > 3 {
			octetLengths[cur] = 1
			cur--
			octetLengths[cur]++
		}
	}

	return result
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("192021"))
	fmt.Println(restoreIpAddresses("0000"))
	fmt.Println(restoreIpAddresses("010010"))
}
