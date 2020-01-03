/**
 * @param {number[]} coins
 * @param {number} amount
 * @return {number}
 */
var coinChange = function (coins, amount) {
  const dp = {};
  const solve = (amount) => {
    if (amount === 0) {
      return 0;
    }

    if (amount < 0) {
      return Number.POSITIVE_INFINITY;
    }

    if (amount in dp) {
      return dp[amount];
    }

    let r = Number.POSITIVE_INFINITY;
    for (const c of coins) {
      if (c <= amount) {
        r = Math.min(r, 1 + solve(amount - c));
      }
    }
    dp[amount] = r;

    return r;
  };

  const r = solve(amount);

  return Number.isFinite(r) ? r : -1;
};

console.log(coinChange([1, 2, 5], 11));
console.log(coinChange([1, 2, 5], 1001));
