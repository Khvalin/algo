/**
 * @param {number} R
 * @param {number} C
 * @param {number} r0
 * @param {number} c0
 * @return {number[][]}
 */
var allCellsDistOrder = function(R, C, r0, c0) {
  let r = []
  for (let x = 0; x < R; x++) {
    for (let y = 0; y < C; y++) {
      r.push({
        d: Math.abs(x - r0) + Math.abs(y - c0),
        p: [ x, y ]
      })
    }
  }

  r = r.sort((a, b) => {
    return a.d - b.d
  })

  const ans = r.map((o) => o.p)

  return ans
}

console.log(allCellsDistOrder(1, 2, 0, 0))
console.log(allCellsDistOrder(2, 3, 1, 2))
