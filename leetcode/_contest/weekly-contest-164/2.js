//count-servers-that-communicate

/**
 * @param {number[][]} grid
 * @return {number}
 */
var countServers = function(grid) {
  const [ rows, cols ] = [ {}, {} ]
  for (let i = 0; i < grid.length; ++i) {
    rows[i] = rows[i] || 0
    for (let j = 0; j < grid[i].length; ++j) {
      cols[j] = cols[j] || 0

      if (grid[i][j]) {
        rows[i]++
        cols[j]++
      }
    }
  }

  let r = 0
  for (let i = 0; i < grid.length; ++i) {
    for (let j = 0; j < grid[i].length; ++j) {
      if (grid[i][j] && (rows[i] > 1 || cols[j] > 1)) {
        r++
      }
    }
  }

  return r
}
