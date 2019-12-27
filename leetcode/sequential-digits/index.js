/**
 * @param {number} low
 * @param {number} high
 * @return {number[]}
 */
var sequentialDigits = function(low, high) {
  if (low > high) {
    return [];
  }

  const l = Math.pow(10, Math.floor(Math.log10(low)));

  let [ n, d, i ] = [ 1, 10, 2 ];
  while (n < l) {
    n = 10 * n + i;
    i++;
    d *= 10;
  }

  const r = [];
  while (i <= 10 && n <= high) {
    if (n >= low) {
      r.push(n);
    }
    n = 10 * n + i;
    n %= d;
    i++;
  }

  r.push(...sequentialDigits(d, high));

  return r;
};

let [ low, high ] = [ 124, 3000 ];

console.log(sequentialDigits(low, high));
