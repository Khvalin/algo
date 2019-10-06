/**
 * @param {number[][]} grid
 * @return {number}
 */
var getMaximumGold = function(grid) {
  const visit = (x, y, visited) => {
    const key = `${x} ${y}`
    if (!grid[x] || !grid[x][y] || visited[key]) {
      return Number.NEGATIVE_INFINITY
    }

    let r = 0
    const t = grid[x][y]

    const coords = [ [ x - 1, y ], [ x + 1, y ], [ x, y - 1 ], [ x, y + 1 ] ]

    for (const pair of coords) {
      const [ x1, y1 ] = pair

      r = Math.max(visit(x1, y1, { ...visited, [key]: true }), r)
    }

    return r + t
  }

  let r = 0
  for (let i = 0; i < grid.length; i++) {
    for (let j = 0; j < grid[i].length; j++) {
      if (grid[i][j]) {
        r = Math.max(r, visit(i, j, {}))
      }
    }
  }

  return r
}
/* 
console.log(getMaximumGold([ [ 0, 6, 0 ], [ 5, 8, 7 ], [ 0, 9, 0 ] ]))
console.log(getMaximumGold([ [ 1, 0, 7 ], [ 2, 0, 6 ], [ 3, 4, 5 ], [ 0, 3, 0 ], [ 9, 0, 20 ] ]))

// prettier-ignore
let g = [[1,0,7,0,0,0],[2,0,6,0,1,0],[3,5,6,7,4,2],[4,3,1,0,2,0],[3,0,5,0,20,0]]
console.log(getMaximumGold(g))
 */
// prettier-ignore
g =[[34,0,1,0,0,0],[0,0,2,0,1,0],[5,4,3,7,4,2],[0,0,5,0,1,4],[0,0,5,0,2,3]]
console.log(getMaximumGold(g))
