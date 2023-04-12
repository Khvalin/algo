/**
 * @param {string} s
 * @return {string}
 */
var originalDigits = function (s) {
  const words = [
    ["four", 4],
    ["five", 5],
    ["seven", 7],
    ["six", 6],
    ["two", 2],
    ["zero", 0],
    ["eight", 8],
    ["one", 1],
    ["nine", 9],
    ["three", 3],
  ];

  const count = {};

  for (const ch of s) {
    count[ch] = count[ch] || 0;
    count[ch]++;
  }

  const ans = words.map(() => 0);
  for (const [w, i] of words) {
    let c = s.length;

    for (const ch of w) {
      c = Math.min(c, count[ch] || 0);
    }

    ans[i] += c;
    for (const ch of w) {
      count[ch] -= c;
    }
  }

  let res = [];
  for (let i = 0; i < ans.length; i++) {
    for (let j = 0; j < ans[i]; j++) {
      res.push(i);
    }
  }

  return res.join("");
};
