/**
 * @param {number[][]} grid
 * @return {number}
 */
var maxDistance = function (grid) {
  const dist = grid.map((row) => row.map(() => Number.POSITIVE_INFINITY));

  let q = [];
  for (let x = 0; x < grid.length; x++) {
    for (let y = 0; y < grid[x].length; y++) {
      if (grid[x][y] === 1) {
        dist[x][y] = 0;
        q.push([x, y]);
      }
    }
  }

  if (q.length === 0 || q.length === grid.length * grid[0].length) {
    return -1;
  }

  let res = 0;
  let i = 0;
  while (i < q.length) {
    for (const len = q.length; i < len; i++) {
      const [x, y] = q[i];
      const d = dist[x][y];
      const dirs = [[x - 1, y], [x + 1, y], [x, y - 1], [x, y + 1]];

      for (const [x1, y1] of dirs) {
        if (
          x1 < 0 || y1 < 0 || x1 >= grid.length || y1 >= grid[x1].length ||
          dist[x1][y1] <= d + 1
        ) {
          continue;
        }

        dist[x1][y1] = d + 1;
        q.push([x1, y1]);
        res = Math.max(res, d + 1);
      }
    }
  }

  return res;
};
