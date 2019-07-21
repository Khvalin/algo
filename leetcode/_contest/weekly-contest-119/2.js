/**
 * @param {number[]} A
 * @return {number}
 */
var largestPerimeter = function(A) {
  let res = 0
  A = A.sort((a, b) => b - a)
  for (let i = 0; i < A.length - 2; i++) {
    let j = i + 1
    let k = j + 1
    if (A[i] + A[j] > A[k] && A[i] + A[k] > A[j] && A[j] + A[k] > A[i]) {
      let p = A[i] + A[j] + A[k]
      if (p > res) {
        res = p
      }
    }
  }

  return res
}

console.log(largestPerimeter([ 3, 2, 3, 4 ]))
