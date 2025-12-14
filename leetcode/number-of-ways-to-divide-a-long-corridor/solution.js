/**
 * @param {string} corridor
 * @return {number}
 */
var numberOfWays = function (corridor) {
  const MOD = 10 ** 9 + 7;
  let s = 0;
  let p = 0;
  let res = 1;

  for (let i = 0; i < corridor.length; i++) {
    const ch = corridor[i];

    if (ch === "P") {
      if (s === 0 || s % 2 === 1) {
        continue;
      }
      p += 1;
    }

    if (ch === "S") {
      s += 1;

      if (s % 2 === 0) {
        res *= p + 1;
        res %= MOD;

        p = 0;
      }
    }

  }

  if (s === 0 || s % 2 === 1) {
    res = 0;
  }

  return res;
};
