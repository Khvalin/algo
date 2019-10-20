/**
 * @param {number[][]} coords
 * @return {boolean}
 */
var checkStraightLine = function(coords) {
  const [ x1, y1 ] = coords[0]
  const [ x2, y2 ] = coords[1]

  const k = (y1 - y2) / (x1 - x2)
  const b = y1 - k * x1

  let cx = 0,
    cy = 0
  for (const pair of coords) {
    const [ x, y ] = pair
    if (x !== x1) {
      cx++
    }

    if (y !== y1) {
      cy++
    }
  }

  if (cx == 0 || cy == 0) {
    return true
  }

  //  console.log(k, b)

  for (const pair of coords) {
    const [ x, y ] = pair

    if (y !== k * x + b) {
      return false
    }
  }

  return true
}

console.log(checkStraightLine([ [ -3, -2 ], [ -1, -2 ], [ 2, -2 ], [ -2, -2 ], [ 0, -2 ] ]))
console.log(checkStraightLine([ [ 0, 1 ], [ 1, 3 ], [ -4, -7 ], [ 5, 11 ] ]))
console.log(
  checkStraightLine([ [ 8, 78 ], [ 2, 18 ], [ -1, -12 ], [ -5, -52 ], [ -4, -42 ], [ -8, -82 ], [ 3, 28 ], [ 9, 88 ] ])
)
