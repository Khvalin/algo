package main

import "fmt"

func treeBottom(tree string) []int {
	s := [2]int{-1, 0}
}

func main() {
	r := treeBottom("(2 (7 (2 () ()) (6 (5 () ()) (11 () ()))) (5 () (9 (4 () ()) ())))")
	fmt.Println(r)
}
