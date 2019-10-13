/**
 * @param {number[][]} queens
 * @param {number[]} king
 * @return {number[][]}
 */
var queensAttacktheKing = function(queens, king) {
  const q = {}

  for (const [ x, y ] of queens) {
    q[`${x}-${y}`] = true
  }

  let r = []

  let x = king[0],
    y = king[1]

  while (x >= 0) {
    x--
    if (q[`${x}-${y}`]) {
      r.push([ x, y ])
      break
    }
  }

  x = king[0]
  y = king[1]

  while (x < 9) {
    x++
    if (q[`${x}-${y}`]) {
      r.push([ x, y ])
      break
    }
  }

  x = king[0]
  y = king[1]

  while (y >= 0) {
    y--
    if (q[`${x}-${y}`]) {
      r.push([ x, y ])
      break
    }
  }

  x = king[0]
  y = king[1]

  while (y < 9) {
    y++
    if (q[`${x}-${y}`]) {
      r.push([ x, y ])
      break
    }
  }

  x = king[0]
  y = king[1]

  while (x < 9 && y < 9) {
    x++
    y++
    if (q[`${x}-${y}`]) {
      r.push([ x, y ])
      break
    }
  }

  x = king[0]
  y = king[1]

  while (x < 9 && y >= 0) {
    x++
    y--
    if (q[`${x}-${y}`]) {
      r.push([ x, y ])
      break
    }
  }

  x = king[0]
  y = king[1]

  while (x >= 0 && y < 9) {
    x--
    y++
    if (q[`${x}-${y}`]) {
      r.push([ x, y ])
      break
    }
  }

  x = king[0]
  y = king[1]

  while (x >= 0 && y >= 0) {
    x--
    y--
    if (q[`${x}-${y}`]) {
      r.push([ x, y ])
      break
    }
  }

  return r
}

let queens = [ [ 0, 1 ], [ 1, 0 ], [ 4, 0 ], [ 0, 4 ], [ 3, 3 ], [ 2, 4 ] ],
  king = [ 0, 0 ]

//prettier-ignore
queens = [[5, 6], [7, 7], [2, 1], [0, 7], [1, 6], [5, 1], [3, 7], [0, 3], [4, 0], [1, 2], [6, 3], [5, 0], [0, 4], [2, 2], [1, 1], [6, 4], [5, 4], [0, 0], [2, 6], [4, 5], [5, 2], [1, 4], [7, 5], [2, 3], [0, 5], [4, 2], [1, 0], [2, 7], [0, 1], [4, 6], [6, 1], [0, 6], [4, 3], [1, 7]], king = [3, 4]

console.log(queensAttacktheKing(queens, king))
