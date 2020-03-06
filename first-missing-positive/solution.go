package main

import "fmt"

func firstMissingPositive(nums []int) int {
  count, min, max := 0, 2, 0
  sum := 0
  for _, n := range nums {
    if n > 0 {
      if min > n {
        min = n
      }
      if max < n {
        max = n
      }

      count++      
      sum += n
    }
  }
  
  if min > 1 {
    return 1
  }
  
  if max == count {
    return max + 1
  }

  prog := ((1 + max) * (1 + count)) >> 1
  
  return prog - sum
}

func main() {
  fmt.Println("2 == ", firstMissingPositive([]int{3,4,-1,1}))
}
