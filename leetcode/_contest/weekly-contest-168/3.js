/**
 * @param {string} s
 * @param {number} maxLetters
 * @param {number} minSize
 * @param {number} maxSize
 * @return {number}
 */
var maxFreq = function(str, maxLetters, minSize, maxSize) {
  let w = {};
  let s = [];
  const candidates = {};
  let res = 0;

  let i = 0;
  for (; i < str.length; i++) {
    const ch = str[i];
    s.push(ch);

    w[ch] = w[ch] || 0;
    w[ch]++;

    while (s.length > minSize || Object.keys(w).length > maxLetters) {
      const f = s.shift();
      w[f]--;
      if (w[f] <= 0) {
        delete w[f];
      }
    }

    if (s.length >= minSize) {
      const c = s.join('');
      candidates[c] = candidates[c] || 0;
      candidates[c]++;

      res = Math.max(candidates[c], res);
    }
  }

  return res;
};

let s = 'aababcaab',
  maxLetters = 2,
  minSize = 3,
  maxSize = 4;

let r = maxFreq(s, maxLetters, minSize, maxSize);
console.log(r);

s = 'aababcaab';
maxLetters = 2;
minSize = 2;
maxSize = 4;
r = maxFreq(s, maxLetters, minSize, maxSize);
console.log(r);

s = 'aaaa';
maxLetters = 1;
minSize = 3;
maxSize = 3;
r = maxFreq(s, maxLetters, minSize, maxSize);
console.log(r);
