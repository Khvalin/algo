package main

import "fmt"

func numRollsToTarget(d int, f int, target int) int {
  M := 1000000007
  if f > target {
    f = target
  }

  memo := make([][]int, 1 + d)
  for i := range memo {
    memo[i] = make([]int, 1 + target)
  }
  
  for i := 1; i <= f; i++ {
    memo[1][i] = 1
  }

  for i := 2; i <= d; i++ {
    for j := 1; j < target; j++ {
      for k := 1; k <= f && k + j <= target; k++ {
        memo[i][k+j] += memo[i - 1][j] 
        memo[i][k+j] %= M
      }
    }
  }

  return memo[d][target]
}

func main() {
  fmt.Println(1, numRollsToTarget(1,6,3))
  fmt.Println(6, numRollsToTarget(2,6,7))
  fmt.Println(222616187, numRollsToTarget(30,30,500))
}
