/**
 * @param {number} n
 * @return {boolean}
 */
var winnerSquareGame = function(n) {
  let memo = new Map();
  memo.set(0, false);

  const squares = [];
  for (let i = 1; i * i <= n; i++) {
    squares.unshift(i * i);
  }

  const solve = (n) => {
    if (memo.has(n)) {
      return memo.get(n);
    }

    let res = false;
    for (const s of squares) {
      if (s <= n && !solve(n - s)) {
        res = true;
      }
    }

    memo.set(n, res);

    return res;
  };

  return solve(n);
};
