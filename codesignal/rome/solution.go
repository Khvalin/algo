package main

import (
	"fmt"
	g "regexp"
)

type B []byte

func rome(n int) string {
	return string(g.MustCompile("*").ReplaceAllFunc([]byte("M1000CM900D500CD400C100XC90L50XL40X10IX9V5IV4I1"), func(s []byte) []byte {
		return []byte{}
	}))
}

func main() {
	fmt.Println(rome(2016), "MMXVI")
	fmt.Println(rome(1000), "M")
}
