/**
 * @param {string} word1
 * @param {string} word2
 * @return {number}
 */
var minDistance = function(word1, word2) {
  const memo = new Map();
  const solve = (pos1, pos2) => {
    const key = pos1 * 501 + pos2;
    if (memo.has(key)) {
      return memo.get(key);
    }

    if (pos1 >= word1.length) {
      return Math.max(0, word2.length - pos2);
    }
    if (pos2 >= word2.length) {
      return Math.max(0, word1.length - pos1);
    }

    let res = null;
    if (word1[pos1] === word2[pos2]) {
      res = solve(pos1 + 1, pos2 + 1);
    }

    if (res === null) {
      res = 1 + Math.min(solve(pos1 + 1, pos2), solve(pos1, pos2 + 1));
    }

    memo.set(key, res);
    return res;
  };

  return solve(0, 0); //, console.log(memo)
};
