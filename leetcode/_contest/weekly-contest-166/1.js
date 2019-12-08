/**
 * @param {number} n
 * @return {number}
 */
var subtractProductAndSum = function(n) {
  let p = 1;
  let s = 0;

  while (n > 0) {
    const d = n % 10;
    p *= d;
    s += d;
    n = Math.floor(n / 10);
  }

  return p - s;
};

console.log(subtractProductAndSum(234));
console.log(subtractProductAndSum(4421));
