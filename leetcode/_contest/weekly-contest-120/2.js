/**
 * @param {number[]} A
 * @return {number}
 */
var maxTurbulenceSize = function(A) {
  let ans = A.length ? 1 : 0

  let [ cur1, cur2 ] = [ ans, ans ]
  for (let i = 1; i < A.length; i++) {
    if (A[i - 1] === A[i]) {
      cur1 = 1
      cur2 = 1
      continue
    }
    if (i % 2 === 0) {
      if (A[i - 1] > A[i]) {
        cur1++
        cur2 = 1
      } else {
        cur2++
        cur1 = 1
      }
    } else {
      if (A[i - 1] < A[i]) {
        cur1++
        cur2 = 1
      } else {
        cur2++
        cur1 = 1
      }
    }

    ans = Math.max(cur1, cur2, ans)
  }

  return ans
}

/* console.log(maxTurbulenceSize([ 9, 4, 2, 10, 7, 8, 8, 1, 9 ]))
console.log(maxTurbulenceSize([ 4, 8, 12, 16 ])) */
console.log(maxTurbulenceSize([ 0, 1, 1, 0, 1, 0, 1, 1, 0, 0 ]))
