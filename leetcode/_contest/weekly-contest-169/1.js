/**
 * @param {number} n
 * @return {number[]}
 */
var sumZero = function(n) {
  const r = [];
  if (n % 2 > 0) {
    r.push(0);
    n--;
  }

  let m = n + 2;
  while (n > 0) {
    r.push(-m);
    r.push(m);
    m--;
    n -= 2;
  }

  return r;
};

let n = 5;
console.log(sumZero(5));
