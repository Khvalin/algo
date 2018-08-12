/**
 * @param {number} R
 * @param {number} C
 * @param {number} r0
 * @param {number} c0
 * @return {number[][]}
 */
var spiralMatrixIII = function(R, C, r0, c0) {
  let r = r0
  let c = c0

  const result = [ [ r, c ] ]

  let count = 1

  let rVec = 1
  let cVec = 0

  let width = 3
  let cur = 0

  while (count < R * C) {
    r += rVec
    c + cVec
    if (r < R && r >= 0 && c < C && c >= 0) {
      result.push([ r, c ])
      count++
      cur++
    }
  }

  return result
}

console.log(spiralMatrixIII(1, 4, 0, 0))
console.log(spiralMatrixIII(5, 6, 1, 4))
