/**
 * @param {number[]} A
 * @return {number}
 */
var subarrayBitwiseORs = function(A) {
  let sumsMap = {}

  let maxSum = []
  let total = 0
  for (let i = A.length - 1; i >= 0; i--) {
    total = total | A[i]
    maxSum[i] = total
  }

  sumsMap[total] = true

  for (let i = 0; i < A.length; i++) {
    sum = A[i]
    sumsMap[sum] = true
    let j = i + 1
    while (sum != maxSum[i] && j < A.length) {
      sum = sum | A[j]
      sumsMap[sum] = true
      j++
    }
  }

  return Object.keys(sumsMap).length
}

console.log(subarrayBitwiseORs([ 1, 1, 2 ]))
console.log(subarrayBitwiseORs([ 1, 2, 4 ]))
