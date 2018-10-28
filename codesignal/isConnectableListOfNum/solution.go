package main

import "fmt"
import "sort"

func isConnectableListOfNum(listOfNum []string) bool {
	sort.Slice(listOfNum, func(i, j int) bool {
		return len(listOfNum[i]) < len(listOfNum[j])
	})
	dic := map[int]bool{}

	for _, w := range listOfNum {
		n := 0
		for i, ch := range w {
			d := int(ch - '0')
			if 0 == d {
				d = 10
			}
			n *= 11
			n += d

			if i == len(w)-1 {
				dic[n] = true
			} else {
				if dic[n] {
					return false
				}
			}
		}
	}

	return true
}

func main() {
	res := isConnectableListOfNum([]string{"911",
		"97625999",
		"91125426"})

	fmt.Println(res)

	res = isConnectableListOfNum([]string{"9999999999",
		"999999999"})

	fmt.Println(res)

}
