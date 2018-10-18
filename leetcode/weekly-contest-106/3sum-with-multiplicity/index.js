/**
 * @param {number[]} A
 * @param {number} target
 * @return {number}
 */
var threeSumMulti = function(A, target) {
  let res = 0
  let sum = {}

  for (let k = 0; k < A.length; k++) {
    const t = target - A[k]
    sum[t] = sum[t] || []
    sum[t].push(k)
  }

  // console.log(sum)

  for (let i = 0; i < A.length; i++) {
    for (let j = i + 1; j < A.length; j++) {
      const t = sum[A[i] + A[j]]
      if (t) {
        console.log(A[i], A[j], target - A[t[0]])

        res += t.length
      }
    }

    res %= 1000000007
  }

  return res
}

//console.log(threeSumMulti([ 1, 1, 2, 2, 3, 3, 4, 4, 5, 5 ], 8))
console.log(threeSumMulti([ 1, 1, 2, 2, 2, 2 ], 5))
