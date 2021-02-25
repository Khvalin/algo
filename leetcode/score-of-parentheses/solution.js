/**
 * @param {string} S
 * @return {number}
 */
var scoreOfParentheses = function(S) {
  const sub = [];
  let d = 0;
  for (let i = 0; i < S.length; i++) {
    sub.push(0);
    const ch = S[i];

    if (ch === '(') {
      d++;
    }
    if (ch === ')') {
      if (S[i - 1] === '(') {
        sub[d]++;
      } else {
        sub[d] += 2 * sub[d + 1];
        sub[d + 1] = 0;
      }

      d--;
    }
  }

  return sub[0] + sub[1];
};
