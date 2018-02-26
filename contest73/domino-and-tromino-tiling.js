/**
 * @param {number} N
 * @return {number}
 */

var numTilings = function(N) {
  const tilings = { 0: 1, 1: 1, 2: 2 };

  const f = function(n) {
    if (tilings[n]) {
      return tilings[n];
    }

    const result = 2 * f(n - 1) + f(n - 3);

    tilings[n] = result;
    return result;
  };

  const result = f(N);

  return result % 1000000007;
};

console.log(numTilings(50));
//312342182
