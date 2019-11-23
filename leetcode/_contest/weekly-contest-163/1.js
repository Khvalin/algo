/**
 * @param {number[][]} grid
 * @param {number} k
 * @return {number[][]}
 */
var shiftGrid = function(grid, k) {
  while (k > 0) {
    let c = []
    for (let i = 0; i < grid.length; i++) {
      c.push(grid[i].pop())
    }

    for (let i = 0; i < c.length; i++) {
      const j = i < c.length - 1 ? i + 1 : 0
      grid[j].unshift(c[i])
    }

    k--
  }

  return grid
}

let g = [ [ 1, 2, 3 ], [ 4, 5, 6 ], [ 7, 8, 9 ] ]
console.log(shiftGrid(g, 1))

let grid = [ [ 3, 8, 1, 9 ], [ 19, 7, 2, 5 ], [ 4, 6, 11, 10 ], [ 12, 0, 21, 13 ] ],
  k = 4
console.log(shiftGrid(grid, k))
