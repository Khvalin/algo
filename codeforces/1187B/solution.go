package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const nmax = 200000 + 1

 var writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

func toInt(buf []byte) (n uint32) {
	for _, v := range buf {
		if v >= '0' && v <= '9' {
			n = n*10 + uint32(v-'0')
		}
	}
	return
}

func solve(m []uint32, t string) uint32 {
	a := [30]uint32{}

	var max uint32

	for _, ch := range t {
		c := ch - 'a'
		a[c]++

		t := (m[a[c]*30+uint32(c)])
		if t > max {
			max = t
		}
	}

	return max
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, nmax)

	reader.ReadLine()

	m := make([]uint32, nmax*30)

	a := [30]uint32{}

	for isPrefix := true; isPrefix; {
		var str []byte
		str, isPrefix, _ = reader.ReadLine()

		for i, ch := range str {
			c := ch - 'a'
			a[c]++

			m[a[c]*30+uint32(c)] = uint32(i)
		}
	}

	bytes, _ := reader.ReadBytes('\n')
	n := toInt(bytes)

	for i := uint32(0); i < n; i++ {
		t, _ := reader.ReadString('\n')

		t = strings.Trim(t, "\r\n")
		if len(t) > 0 {
			r := solve(m, t) + 1
			
			printf("%d\n",r)
		}
	}

		writer.Flush()
}
