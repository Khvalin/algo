package main

import (
	"bufio"
	"fmt"
	"os"
)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

func toInt(buf []byte) (n uint32) {
	for _, v := range buf {
		n = n*10 + uint32(v-'0')
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

/*
type index struct {
	ch    rune
	count uint32
}
*/
func main() {
	defer writer.Flush()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	toInt(scanner.Bytes())

	m := make([]uint32, 200000*30)

	scanner.Scan()
	str := scanner.Text()
	a := [30]uint32{}

	for i, ch := range str {
		c := ch - 'a'
		a[c]++

		m[a[c]*30+uint32(c)] = uint32(i)
	}

	scanner.Scan()
	n := toInt(scanner.Bytes())

	for i := uint32(0); i < n; i++ {
		scanner.Scan()
		t := scanner.Text()

		r := solve(m, t) + 1
		printf("%d\n", r)
	}
}
