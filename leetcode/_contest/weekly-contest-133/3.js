/**
 * @param {number[]} A
 * @param {number} L
 * @param {number} M
 * @return {number}
 */
var maxSumTwoNoOverlap = function(A, L, M) {
  const sums1 = [],
    sums2 = []

  let [ s1, s2 ] = [ 0, 0 ]
  for (let i = 0; i < A.length; i++) {
    sums1.push(null)
    sums2.push(null)

    let n = A[i]
    s1 += n
    s2 += n

    if (i >= L - 1) {
      if (i > L - 1) {
        s1 -= A[i - L]
      }
      sums1[i] = s1
    }

    if (i >= M - 1) {
      if (i > M - 1) {
        s2 -= A[i - M]
      }
      sums2[i] = s2
    }
  }

  let r = Number.NEGATIVE_INFINITY
  for (let i = 0; i < A.length; i++) {
    if (sums1[i] === null) {
      continue
    }

    for (let j = 0; j < A.length; j++) {
      if (sums2[j] !== null && (j <= i - L || j >= i + M)) {
        r = Math.max(r, sums1[i] + sums2[j])
      }
    }
  }

  return r
}

let r
r = maxSumTwoNoOverlap([ 0, 6, 5, 2, 2, 5, 1, 9, 4 ], 1, 2)
console.log(r)

r = maxSumTwoNoOverlap([ 3, 8, 1, 3, 2, 1, 8, 9, 0 ], 3, 2)
console.log(r)

r = maxSumTwoNoOverlap([ 2, 1, 5, 6, 0, 9, 5, 0, 3, 8 ], 4, 3)
console.log(r)
