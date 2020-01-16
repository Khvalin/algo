package main

import (
	"bufio"
	"fmt"
	"os"
)

const NMAX = 100
const EMPTY = 255
const INFINITY = uint32(999999)

func toInt(buf []byte) uint32 {
	n := uint32(0)
	for _, v := range buf {
		n = n*10 + uint32(v-'0')
	}

	return n
}

func solve(n uint32, p [NMAX]uint32) uint32 {
  type State struct  {
    odds int
    evens int
    last uint32
    index uint32
  }
  
  start := State{}
  
  for i := uint32(1); i <= n; i++ {
    f := false
    for j := uint32(0); j < n; j++ {
      f = f || i == p[j]
    }

    if !f {
      if i % 2 == 0 {
        start.evens++
      } else {
        start.odds++
      }
    }
  }
  
  memo := map[State]uint32{}
  var solve func(s State) uint32
  solve = func(s State) uint32 {
    if r, f := memo[s]; f {
      return r
    }

    i, r := s.index, uint32(0)
    for ; i < n && p[i] != 0; i++ {
      if i > 0 {
        prev := p[i - 1]
        if prev == 0 {
          prev = s.last
        }
        if prev % 2 != p[i] % 2 {
          r++
        }
      }
    }
    
    prev := uint32(0)
    if i > 0 {
      prev = p[i - 1]
      if prev == 0 {
        prev = s.last
      }
    }
    
    if i < n && p[i] == 0 {
      r1, r2 := INFINITY, INFINITY
      if s.odds > 0 {
        next := State{index: i + 1, odds: s.odds - 1, evens: s.evens, last: 1}
        r1 = solve(next)
        if i > 0 && (prev % 2 == 0) {
          r1++
        }
      }
      
      if s.evens > 0 {
        next := State{index: i + 1, odds: s.odds, evens: s.evens - 1}
        r2 = solve(next)
        if i > 0 && (prev % 2 > 0) {
          r2++
        }
      }

      if r1 > r2 {
        r1 = r2
      }
      
      r += r1
    }
    
    memo[s] = r
    
    return r
  }

  return solve(start)
}

func main() {
	var n uint32

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n = toInt(scanner.Bytes())
	
  p := [NMAX]uint32{}
  for i := uint32(0); i < n; i++ {
    scanner.Scan()
    p[i] = toInt(scanner.Bytes())
  }
  
  res := solve(n, p)
  fmt.Println(res)
}
