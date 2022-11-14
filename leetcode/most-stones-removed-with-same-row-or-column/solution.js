/**
 * @param {number[][]} stones
 * @return {number}
 */
var removeStones = function (stones) {
  const visited = stones.map(() => false);
  const comp = stones.map((_, i) => i);

  const walk = (ind, mark) => {
    if (visited[ind]) {
      return;
    }

    visited[ind] = true;
    comp[ind] = mark;
    for (let i = 0; i < stones.length; i++) {
      if (
        !visited[i] &&
        (stones[ind][0] === stones[i][0] || stones[ind][1] === stones[i][1])
      ) {
        walk(i, mark);
      }
    }
  };

  for (let i = 0; i < stones.length; i++) {
    if (!visited[i]) {
      walk(i, i);
    }
  }

  let res = stones.length;
  for (let i = 0; i < stones.length; i++) {
    if (comp[i] === i) {
      res--;
    }
  }

  return res;
};
