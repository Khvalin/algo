package main

import (
	"fmt"
	"math"
	"math/big"
	"strings"
)

func extraCreditAssignment(decimal string, denominator int) int {
	l := strings.IndexRune(decimal, '(')
	r := strings.IndexRune(decimal, ')')

	if l > -1 {
		rep := decimal[l+1 : r]
		c := 20000 / len(rep)
		if c < 2 {
			c = 4
		}
		rep = strings.Repeat(rep, c)
		decimal = decimal[:l] + rep + decimal[r+1:]
	}

	f := new(big.Float)
	_, err := fmt.Sscan(decimal, f)

	res := float64(-1)
	if err == nil {
		z := new(big.Float).Mul(f, big.NewFloat(float64(denominator)))
		res, _ = z.Float64()
	}

	return int(math.Round(res))
}

func main() {
	fmt.Println(extraCreditAssignment("123456.(13)", 99))
	fmt.Println(extraCreditAssignment("5.(717948)", 78))
	fmt.Println(extraCreditAssignment("3.(142857)", 7))
	fmt.Println(extraCreditAssignment("0.(008403361344537815126050420168067226890756302521)", 199))
}
