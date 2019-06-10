/**
 * @param {character[][]} grid
 * @return {number}
 */
var numIslands = function(grid) {
  const map = []
  for (let i = 0; i < grid.length; i++) {
    map.push([])
    for (let j = 0; j < grid[i].length; j++) {
      map[i].push(false)
      grid[i][j] = 1 * grid[i][j]
    }
  }
  const dfs = (x, y) => {
    if (map[x][y] || !grid[x][y]) {
      return
    }
    map[x][y] = true
    dfs(Math.min(x + 1, grid.length - 1), y)
    dfs(x, Math.min(y + 1, grid[x].length - 1))

    dfs(Math.max(x - 1, 0), y)
    dfs(x, Math.max(y - 1), 0)
  }

  let res = 0
  for (let i = 0; i < grid.length; i++) {
    for (let j = 0; j < grid[i].length; j++) {
      if (grid[i][j] > 0 && !map[i][j]) {
        dfs(i, j)
        res++
      }
    }
  }

  return res
}

console.log(
  numIslands([
    [ '1', '1', '0', '0', '0' ],
    [ '1', '1', '0', '0', '0' ],
    [ '0', '0', '1', '0', '0' ],
    [ '0', '0', '0', '1', '1' ]
  ])
)
