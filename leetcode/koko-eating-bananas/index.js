//https://leetcode.com/problems/koko-eating-bananas/
// binary search

/**
 * @param {number[]} piles
 * @param {number} H
 * @return {number}
 */
var minEatingSpeed = function(piles, H) {
    let hi = Math.max(...piles)
    let lo = 1

    while (lo < hi) {
      let m = (hi + lo) >> 1

      let s = 0
      for (const n of piles) {
        s += Math.ceil(n/m)
      }

      if (s > H) {
        lo = m + 1
      } else {
        hi = m
      }
    }

  return lo
};
