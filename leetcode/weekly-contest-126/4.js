/**
 * @param {number[]} stones
 * @param {number} K
 * @return {number}
 */
var mergeStones = function(stones, K) {
  const MAX = Number.POSITIVE_INFINITY
  const memo = {}

  const solve = function(stones) {
    console.log(stones)
    if (stones.length == 1) {
      return 0
    }

    if (stones.length < K) {
      return MAX
    }

    const key = stones.join(' ')

    if (memo[key]) {
      return memo[key]
    }

    let min = MAX

    let cur = 0
    // let minIndex = -1

    for (let i = 0; i < stones.length; i++) {
      cur += stones[i]

      if (i >= K) {
        cur -= stones[i - K]
      }

      if (i >= K - 1) {
        let r = cur + solve([ ...stones.slice(0, i - K + 1), cur, ...stones.slice(i + 1) ])
        if (min > r) {
          min = r
        }
      }
    }
    // }

    memo[key] = min
    return min
  }

  let res = solve(stones, K)
  if (res == MAX) {
    res = -1
  }

  return res
}
/* 
console.log(mergeStones([ 3, 2, 4, 1 ], 2))

console.log(mergeStones([ 3, 2, 4, 1 ], 3))
console.log(mergeStones([ 1, 1, 1, 1, 100 ], 2))
 */
console.log(
  mergeStones([ 46, 96, 43, 76, 38, 38, 51, 86, 5, 19, 30, 73, 66, 2, 55, 23, 32, 13, 86, 100, 95, 24, 17, 40, 14 ], 2)
)
