/**
 * @param {number[]} A
 * @param {number} K
 * @return {number}
 */
var subarraysDivByK = function(A, K) {
  let res = 0
  let memo = Array.from(A)

  for (let i = 0; i < A.length; i++) {
    //	copy(memo1, memo2)

    for (let j = 0; j < A.length - i; j++) {
      if (i == 0) {
        memo[j] = 0
      }
      memo[j] += A[i + j]

      //  console.log(memo[0])

      if (memo[j] % K == 0) {
        res++
      }
    }
  }

  return res
}

console.log(subarraysDivByK([ 4, 5, 0, -2, -3, 1 ], 5))
