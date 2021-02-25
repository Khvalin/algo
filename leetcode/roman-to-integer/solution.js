/**
 * @param {string} s
 * @return {number}
 */
var romanToInt = function(s) {
  const R = { I: 1, V: 5, X: 10, L: 50, C: 100, D: 500, M: 1000 };

  let res = 0;
  let i = 0;
  while (i < s.length) {
    let n = R[s[i]];
    i++;

    if (i < s.length && n < R[s[i]]) {
      n = R[s[i]] - n;
      i++;
    }

    res += n;
  }

  return res;
};
