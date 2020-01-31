package main

import "fmt"

func findNumberOfLIS(nums []int) int {
  l := len(nums)
  memo := make([][]int, l)
  for i := range memo {
    memo[i] = make([]int, 1 + l)
    memo[i][1] = 1
  }

  count := make([]int, l)
  
  var solve func(ind int) int
  solve = func(ind int) int {
    if ind >= l {
      return 0
    }
    
    if count[ind] > 0 {
      return count[ind]
    }
    
    m := 1

    for i := ind + 1; i < l; i++ {
      if nums[ind] < nums[i] {
        t := 1 + solve(i)

        memo[ind][t] += memo[i][t - 1]
        if t > m {
          m = t
        }
      }
    }
    
    count[ind] = m

    return m
  }
  
  m := 0
  for i := range nums {
    if t := solve(i); m < t {
      m = t
    }
  }
  
  r := 0
  for i, n := range count {
    if n == m {
      r += memo[i][n]
    }
  }

  return r
}


func main() {
  fmt.Println(findNumberOfLIS( []int {2,2,2,2,2} ))
  fmt.Println(findNumberOfLIS( []int {12, 1, 2, 4, 3, 5} ))
}
