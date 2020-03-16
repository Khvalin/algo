package main

import "fmt"

func booleanParenthesization(expression string) int {
  var mod uint = 1003
  /*E := map[string] [2]uint {
    "": [2]uint{0, 0},
    "T&T": [2]uint{0, 1},
    "T|T": [2]uint{0, 1},
    "F|T": [2]uint{0, 1},
    "T|F": [2]uint{0, 1},
    "T^F": [2]uint{0, 1},
    "F^T": [2]uint{0, 1},
  }*/

  memo := map[int][2]uint{}

  
  var count func(l, r int) [2]uint
  count = func(l, r int) [2]uint {
    key := l * 1e3 +r
    res, f:= memo[key]
    
    if f {
      return res
    }
    
    if l >= r {
      return res
    }
    
    if r-l== 1 {
      if expression[l] == 'T' {
        res[1]++
      } else {
        res[0]++
      }
      
      return res
    }

    for i := l+1; i < r; i+=2 {
      lcount, rcount := count(l, i), count(i + 1, r)
      switch expression[i] {
        case '&':
          min := lcount[1]
          if rcount[1] < min {
            min = rcount[1]
          }
          res[0] += lcount[0] + rcount[0] + lcount[1] + rcount[1] - min
          res[1] += min
        break
        case '|':
          min := lcount[0]
          if rcount[0] < min {
            min = rcount[0]
          }
          res[0] += min
          res[1] += lcount[0] + rcount[0] + lcount[1] + rcount[1] - min
        break
        case '^':
          res[0] += rcount[0] * lcount[0] + rcount[1] * lcount[1] 
          res[1] += rcount[0] * lcount[1] + rcount[1] * lcount[0] 
        break
      }
      
      res[0] %= mod
      res[1] %= mod
    }
    
    memo[key] = res
    
    return res
  }
  
  res := count(0, len(expression))

  return int(res[1] % mod)
}

func main() {
	fmt.Println(5, booleanParenthesization("T&T|F^F"))
}
