/**
 * @param {number[][]} grid
 * @return {number}
 */
var uniquePathsIII = function(grid) {
  let start, end
  let N = 0
  let res = 0

  const dfs = (p, visited = {}) => {
    let key = (p.x + 1) * 21 + p.y
    if (!grid[p.x] || !Number.isInteger(grid[p.x][p.y]) || grid[p.x][p.y] == -1 || visited[key]) {
      return
    }
    if (p.x === end.x && p.y === end.y && p.c === N + 1) {
      // console.log(p)
      res++
      return
    }

    // console.log(p)

    //if (grid[p.x][p.y] == 0) {
    dfs({ ...p, x: p.x - 1, c: p.c + 1 }, { ...visited, [key]: true })
    dfs({ ...p, x: p.x + 1, c: p.c + 1 }, { ...visited, [key]: true })
    dfs({ ...p, y: p.y - 1, c: p.c + 1 }, { ...visited, [key]: true })
    dfs({ ...p, y: p.y + 1, c: p.c + 1 }, { ...visited, [key]: true })
    // }
  }

  for (let i = 0; i < grid.length; i++) {
    for (let j = 0; j < grid[i].length; j++) {
      if (grid[i][j] === 0) {
        N++
      }
      if (grid[i][j] === 1) {
        start = { x: i, y: j, c: 0 }
      }

      if (grid[i][j] === 2) {
        end = { x: i, y: j }
      }
    }
  }

  dfs(start)
  return res
}

console.log(uniquePathsIII([ [ 1, 0, 0, 0 ], [ 0, 0, 0, 0 ], [ 0, 0, 2, -1 ] ]))
