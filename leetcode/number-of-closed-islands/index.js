/**
 * @param {number[][]} grid
 * @return {number}
 */
var closedIsland = function(grid) {
  const n = grid.length
  const m = grid[0].length

  const visited = []
  for (let i = 0; i < n; i++) {
    visited.push(new Array(m))
    visited[i].fill(false, 0)
  }

  const visit = (x, y) => {
    const cells = [ [ x, y ] ]
    let res = false

    while (cells.length) {
      const [ x, y ] = cells.shift()
      if (visited[x][y] || grid[x][y] != 0) {
        continue
      }
      visited[x][y] = true

      const next = [ [ x + 1, y ], [ x - 1, y ], [ x, y + 1 ], [ x, y - 1 ] ]
      for (const pair of next) {
        const [ x, y ] = pair

        if (x < 0 || y < 0 || x >= n || y >= m) {
          res = true
          continue
        }

        if (grid[x][y] == 0 && !visited[x][y]) {
          cells.push(pair)
        }
      }
    }

    return res
  }

  let r = 0
  for (let i = 0; i < n; i++) {
    for (let j = 0; j < m; j++) {
      if (grid[i][j] === 0 && !visited[i][j]) {
        if (!visit(i, j)) {
          r++
        }
      }
    }
  }

  return r
}

let grid = [
  [ 1, 1, 1, 1, 1, 1, 1, 0 ],
  [ 1, 0, 0, 0, 0, 1, 1, 0 ],
  [ 1, 0, 1, 0, 1, 1, 1, 0 ],
  [ 1, 0, 0, 0, 0, 1, 0, 1 ],
  [ 1, 1, 1, 1, 1, 1, 1, 0 ]
]
console.log(closedIsland(grid))

grid = [ [ 0, 0, 1, 0, 0 ], [ 0, 1, 0, 1, 0 ], [ 0, 1, 1, 1, 0 ] ]
console.log(closedIsland(grid))
