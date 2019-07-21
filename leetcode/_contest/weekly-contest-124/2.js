/**
 * @param {number[][]} grid
 * @return {number}
 */
var orangesRotting = function(grid) {
  let freshCount = 0
  let rotten = []
  let total = 0
  for (let i = 0; i < grid.length; i++) {
    for (let j = 0; j < grid[i].length; j++) {
      if (grid[i][j] === 1) {
        freshCount++
        total++
      }
      if (grid[i][j] === 2) {
        rotten.push({ x: i, y: j })
        total++
      }
    }
  }

  let prev = -1
  let newRotten = rotten.length
  let res = 0

  while (freshCount > 0 && newRotten > 0) {
    res++
    prev = newRotten
    const todo = []
    for (const rot of rotten) {
      if (grid[rot.x - 1]) {
        if (grid[rot.x - 1][rot.y] == 1) {
          grid[rot.x - 1][rot.y] = 2

          todo.push({ x: rot.x - 1, y: rot.y })
        }
      }

      if (grid[rot.x + 1]) {
        if (grid[rot.x + 1][rot.y] == 1) {
          grid[rot.x + 1][rot.y] = 2
          todo.push({ x: rot.x + 1, y: rot.y })
        }
      }

      if (grid[rot.x][rot.y - 1] == 1) {
        grid[rot.x][rot.y - 1] = 2
        todo.push({ x: rot.x, y: rot.y - 1 })
      }

      if (grid[rot.x][rot.y + 1] == 1) {
        grid[rot.x][rot.y + 1] = 2
        todo.push({ x: rot.x, y: rot.y + 1 })
      }
    }

    newRotten = todo.length
    freshCount -= newRotten

    rotten = [ ...todo ]

    //console.log(newRotten, freshCount, rotten)
  }

  if (freshCount > 0) {
    res = -1
  }

  return res
}

console.log(orangesRotting([ [ 2, 1, 1 ], [ 1, 1, 0 ], [ 0, 1, 1 ] ]))
console.log(orangesRotting([ [ 2, 1, 1 ], [ 0, 1, 1 ], [ 1, 0, 1 ] ]))
console.log(orangesRotting([ [ 0, 2 ] ]))
