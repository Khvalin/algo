/**
 * @param {number[]} arr
 * @param {number} difference
 * @return {number}
 */
var longestSubsequence = function(arr, d) {
  let m = {}
  let r = 0

  for (const n of arr) {
    m[n] = Math.max((m[n - d] || 0) + 1, m[n] || 1)

    r = Math.max(r, m[n])
  }

  return r
}

console.log(longestSubsequence([ 1, 2, 3, 4 ], 1))
console.log(longestSubsequence([ 1, 3, 5, 7 ], 1))
// prettier-ignore
console.log(longestSubsequence( [1,5,7,8,5,3,4,2,1],  -2))

// prettier-ignore
console.log(longestSubsequence([4,12,10,0,-2,7,-8,9,-9,-12,-12,8,8], 0))

console.log(longestSubsequence([ 2, -6, -3, -6, 2, 0 ], -2))
