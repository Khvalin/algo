package main

import (
	"fmt"
)

func marvelousNumber(n string) string {
	dgts, L := []byte(n), len(n)

	for i := range dgts {
		dgts[i] -= '0'
	}

	first := L + 1
	for i := 0; i < L; i++ {
		if first > i && dgts[first] > 7 {
			first = i
		}
	}

	return ""
}

func main() {
	fmt.Println(marvelousNumber("4587"))
	fmt.Println(marvelousNumber("1007"))
}
