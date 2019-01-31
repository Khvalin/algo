func arrayMaxConsecutiveSum(inputArray []int, k int) int {
  s := 0
  r := inputArray[0]
  for i, n := range inputArray {
    s += n
    if i >= k {
      s -= inputArray[i-k]
    }
    if s > r {
      r = s
    }
  }
  
  return r
}
