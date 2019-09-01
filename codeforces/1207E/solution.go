package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100

func write(a string) {
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()

	f.Write([]byte(a))
	f.Write([]byte("\n"))
}

func gen() [N]uint64 {
	res := [N]uint64{}
	for i := uint64(1); i <= N; i++ {
		res[i-1] = i
	}

	return res
}

func join(a [N]uint64) string {
	s := fmt.Sprint(a)
	s = s[1 : len(s)-1]

	return s
}

func main() {
	ans := [2]uint64{}

	for j := 0; j < 2; j++ {
		sample := gen()

		for i, n := range sample {
			if j == 1 {
				sample[i] = n << 7
			}
		}

		q := "? " + join(sample)
		write(q)

		fmt.Scan(&ans[j])
	}

	res := (ans[0]>>7)<<7 + ans[1]&127
	fmt.Println("!", res)
}
