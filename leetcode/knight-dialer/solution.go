package main

import "fmt"

func knightDialer(n int) int {
  const nmax = uint(1000000007)
  adj := [10][3]int{
    [3]int{4, 6, -1}, //0
    [3]int{6, 8, -1},
    [3]int{7, 9, -1}, // 2
    [3]int{4, 8, -1}, // 3
    [3]int{0, 3, 9 }, // 4
    [3]int{-1}, // 5
    [3]int{0, 1, 7}, // 6
    [3]int{2, 6, -1}, // 7
    [3]int{1, 3, -1}, // 8
    [3]int{2, 4, -1}, // 9
  }
  
  memo := [2][10]uint{}
  for i := 0; i < 10; i++ {
    memo[1][i] = uint(1)
  }

  for i := 2; i <= n; i++ {
    cur, prev := i % 2, (i+1)%2
    for j := 0; j < 10; j++ {
      sum := uint(0)
      for k := 0; k < len(adj[prev]) && adj[j][k] > -1; k++ {
        sum += memo[prev][ (adj[j][k]) ]
      }
      
      memo[cur][j] = sum % nmax
    }
  }
  
  r := uint(0)
  for i := 0; i < 10; i++ {
    r += memo[n%2][i]
    r %= nmax
  }
  
  return int(r)
}

func main() {
  fmt.Println("10 == ", knightDialer(1))
  fmt.Println("20 == ", knightDialer(2))
  fmt.Println("46 == ", knightDialer(3))
  fmt.Println("406880451 == ", knightDialer(5000))
}
