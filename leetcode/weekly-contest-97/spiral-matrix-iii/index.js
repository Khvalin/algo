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

  let rVec = 0
  let cVec = 1

  let width = 1

  while (count < R * C) {
    for (j = 0; j < 2; j++) {
      for (let i = 0; i < width; i++) {
        r += rVec
        c += cVec

        //console.log([ r, c ])
        if (r < R && r >= 0 && c < C && c >= 0) {
          result.push([ r, c ])
          count++
        }
      }

      if (rVec == 0) {
        rVec = cVec
        cVec = 0
      } else {
        cVec = -rVec
        rVec = 0
      }
    }
    width += 1
  }

  return result
}

console.log(spiralMatrixIII(1, 4, 0, 0))
console.log(spiralMatrixIII(5, 6, 1, 4))
