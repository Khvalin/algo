package main

import "fmt"

// The value of each non-leaf node is equal to the product of the largest leaf value in its left and right subtree respectively.
func mctFromLeafValues(arr []int) int {
  const inf = 2147483649

  type state struct {
    left, right int
  }
  
  type tree struct {
    max, sum int
  }
  
  pows := make([]bool, len(arr) + 2)
  for i := 2; i < len(pows); i *= 2 {
    pows[i] = true
  }

  memo := map[state]tree{}
  for i := 0; i < len(arr) - 1;i++ {
    val := arr[i] * arr[i + 1]
    
    t := tree{max: arr[i], sum: val}
    if t.max < arr[i + 1] {
      t.max = arr[i + 1]
    }
    memo[state{i, i + 1}] = t
  }

  var solve func(left, right int) tree
  solve = func(left, right int) tree {
    st := state{left, right}

    if res, found := memo[st]; found {
      return res
    }

    leftTree := solve(left + 1, right)
    leftTree.sum += arr[left] * leftTree.max
    if leftTree.max < arr[left] {
      leftTree.max = arr[left]
    }

    rightTree := solve(left, right - 1)
    rightTree.sum += arr[right] * leftTree.max
    if rightTree.max < arr[right] {
      rightTree.max = arr[right]
    }
    
    midTree := tree{sum: inf}
    l := right - left + 1
    if pows[l] {
      lHalf := solve(left, left + l >> 1 - 1)
      rHalf := solve(left + l >> 1, right)

      midTree.sum = lHalf.sum + rHalf.sum
      midTree.sum += lHalf.max * rHalf.max
      midTree.max = lHalf.max
      if lHalf.max < rHalf.max {
        midTree.max = rHalf.max
      }
    }

    res := midTree
    if res.sum > leftTree.sum {
      res = leftTree
    }
    
    if res.sum > rightTree.sum {
      res = rightTree
    }
    
    memo[st] = res

    return res
  }

  tr := solve(1, len(arr) - 1)

  return tr.sum
}


func main() {
//  fmt.Println(32, mctFromLeafValues([]int{6,2,4}))
//  fmt.Println(15, mctFromLeafValues([]int{1, 3, 3, 1}))
  fmt.Println(541, mctFromLeafValues([]int{15,10, 2, 3, 4, 1, 1, 1, 5, 4, 3, 6, 15}))
}
