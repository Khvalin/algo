/**
 * @param {number[]} A
 * @param {number} K
 * @return {number}
 */
var largestSumAfterKNegations = function(A, K) {
  let a = [ ...A ]

  let i = 0

  //let sorted = false
  for (let j = 0; j < K; j++) {
    a = a.sort((a, b) => a - b)
    a[0] = -a[0]
  }

  return a.reduce((acc, n) => {
    acc += n
    return acc
  }, 0)
}

largestSumAfterKNegations([ 4, 2, 3 ], 1)
