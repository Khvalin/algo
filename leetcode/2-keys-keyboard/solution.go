func minSteps(n int) int {
  memo := map[int]int{1: 0}

  var solve func(n int) int
  solve = func(n int) int {
    if r, f := memo[n]; f {
      return r
    }
      
    r := 9999
    for i := 1; i <= n >> 1; i++ {
      if n % i == 0 {
        d := n / i
        d += solve(i)
        if d < r {
          r = d
        }
      }
    }
    
    memo[n] = r
    
    return r
  }
  
  return solve(n)
}