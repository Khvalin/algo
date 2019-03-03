/**
 * @param {number[]} A
 * @param {number} K
 * @return {number}
 */
var longestOnes = function(A, K) {
  let res = 0
  let cur = 0
  let k = K
  let flips = []
  for (let i = 0; i < A.length; i++) {
    let d = A[i]
    if (d === 0) {
      if (k > 0) {
        k--
        d = 1
        flips.push(i)
      } else {
        if (K > 0) {
          let first = flips.shift()
          cur = i - first - 1
          flips.push(i)
          d = 1
        }
      }
    }
    if (d === 1) {
      cur++
    }
    res = Math.max(res, cur)
    //
  }

  return res
}

console.log(longestOnes([ 0, 0, 0, 1, 1, 1, 1, 0 ], 2))
console.log(longestOnes([ 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0 ], 2))
console.log(longestOnes([ 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1 ], 3))

// prettier-ignore
let a=[1,0,0,0,1,1,0,0,1,1,0,0,0,0,0,0,1,1,1,1,0,1,0,1,1,1,1,1,1,0,1,0,1,0,0,1,1,0,1,1]
let K = 8

console.log(longestOnes(a, K))
