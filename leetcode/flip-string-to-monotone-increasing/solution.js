/**
 * @param {string} s
 * @return {number}
 */
var minFlipsMonoIncr = function (s) {
  const count = [new Array(s.length).fill(0), new Array(s.length).fill(0)];

  // left to right
  for (let i = 0; i < s.length; i++) {
    if (i > 0) {
      count[0][i] = count[0][i - 1];
      count[1][i] = count[1][i - 1];
    }

    const ind = s[i] === "0" ? 1 : 0;
    count[ind][i]++;
  }

  let res = s.length + 10;

  // right to left
  let c = [0, 0];
  for (let i = s.length - 1; i >= 0; i--) {
    if (s[i] === "0") {
      c[1]++;
    } else {
      c[0]++;
    }
    if (i > 0) {
      res = Math.min(count[0][i - 1] + c[1], count[1][i - 1] + c[1], res);
    }
  }

  res = Math.min(...c, res);

  return res;
};

console.log(minFlipsMonoIncr("10011111110010111011"));
console.log(minFlipsMonoIncr("1111001110"));
console.log(minFlipsMonoIncr("100000001010000"));
