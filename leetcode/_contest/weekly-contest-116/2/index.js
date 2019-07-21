/**
 * @param {number[]} A
 * @return {number}
 */
var maxWidthRamp = function(A) {
  let arr = A.map((val, index) => ({ val, index }))
  arr = arr.sort((a, b) => a.val - b.val)

  let maxDist = 0
  for (const tuple of arr) {
    for (let i = A.length - 1; i > tuple.index; i--) {
      if (A[i] >= tuple.val) {
        maxDist = Math.max(maxDist, i - tuple.index)
        break
      }
    }
  }

  return maxDist
}

console.log(maxWidthRamp([ 6, 0, 8, 2, 1, 5 ]))
console.log(maxWidthRamp([ 9, 8, 1, 0, 1, 9, 4, 0, 4, 1 ]))
