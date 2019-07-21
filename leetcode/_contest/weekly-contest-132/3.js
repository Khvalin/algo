/**
 * @param {number[]} A
 * @return {number}
 */
var longestArithSeqLength = function(A) {
  let res = 0
  let sums = {}

  for (let i = 0; i < A.length; i++) {
    for (let j = i + 1; j < A.length; j++) {
      let d = A[i] - A[j]
      sums[d] = sums[d] || []
      sums[d].push([ i, j ])
    }
  }

  for (const s of Object.values(sums)) {
    while (s.length > 0) {
      const p = s.shift()
      let n = p[1]

      const nums = new Set(p)

      for (const pair of s) {
        if (pair[0] === n) {
          n = pair[1]

          nums.add(pair[0])
          nums.add(pair[1])
        }
      }

      res = Math.max(res, nums.size)
    }
  }

  return res
}

console.log(longestArithSeqLength([ 3, 6, 9, 12 ]))

console.log(longestArithSeqLength([ 20, 1, 15, 3, 10, 5, 8 ]))
console.log(longestArithSeqLength([ 9, 4, 7, 2, 10 ]))
console.log(longestArithSeqLength([ 83, 20, 17, 43, 52, 78, 68, 45 ]))
