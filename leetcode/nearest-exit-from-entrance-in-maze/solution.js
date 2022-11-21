/**
 * @param {character[][]} maze
 * @param {number[]} entrance
 * @return {number}
 */
var nearestExit = function (maze, entrance) {
  const a = [entrance];
  const visited = new Set();
  visited.add(101 * entrance[0] + entrance[1]);

  let d = 0;
  let i = 0;
  while (i < a.length) {
    d++;

    const len = a.length;
    for (; i < len; i++) {
      const [x, y] = a[i];
      const moves = [
        [x - 1, y],
        [x + 1, y],
        [x, y - 1],
        [x, y + 1],
      ];

      for (const [dx, dy] of moves) {
        const key = 101 * dx + dy;
        if (
          dx < 0 ||
          dx >= maze.length ||
          dy < 0 ||
          dy >= maze[0].length ||
          maze[dx][dy] !== "." ||
          visited.has(key)
        ) {
          continue;
        }

        if (dx === 0 || dy === 0 || dx === maze.length - 1 || dy === maze[0].length - 1) {
            return d
        }

        a.push([dx, dy])
        visited.add(key)
      }
    }
  }

  return -1;
};
