/**
 * @param {number[]} A
 * @return {boolean}
 */
var validMountainArray = function(A) {
  let N = A.length
  if (N < 3) {
    return false
  }

  let i = 0
  while (A[i + 1] > A[i]) {
    i++
  }

  if (i == 0 || i >= N - 1) {
    return false
  }
  i++

  while (A[i] < A[i - 1]) {
    i++
  }

  return i == N
}

console.log(validMountainArray([ 9, 8, 7, 6, 5, 4, 3, 2, 1, 0 ]))
console.log(validMountainArray([ 1, 2, 3 ]))
