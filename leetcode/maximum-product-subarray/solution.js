/**
 * @param {number[]} nums
 * @return {number}
 */
var maxProduct = function(nums) {
  const solve = (left, right) => {
    let res = nums[left];

    let m = null;
    let neg = null;
    for (let i = left; i < right; i++) {
      const n = nums[i];
      res = Math.max(n, res);
      if (n <= 0) {
        if (n < 0) {
          if (neg === null) {
            neg = n;
          } else {
            neg *= n;
            res = Math.max(neg, res);
          }
        }

        if (m !== null) {
          res = Math.max(m, res);
        }

        m = null;
        continue;
      }

      if (m === null) {
        m = n;
      } else {
        m *= n;
      }
      res = Math.max(m, res);

      if (neg === null) {
        neg = n;
        continue;
      }

      neg *= n;
      res = Math.max(neg, res);
    }

    if (neg < 0) {
      let i = left;
      while (i < right - 1 && neg < 0) {
        neg /= nums[i];
        i++;
      }
    }

    return Math.max(m || nums[0], neg || nums[0], res);
  };

  let left = 0;
  const findZero = (n, i) => {
    return n === 0 && i >= left;
  };

  let res = Math.max(...nums);

  do {
    let right = nums.findIndex(findZero);
    if (right < 0) {
      right = nums.length;
    }

    res = Math.max(res, solve(left, right) || nums[0]);

    left = right + 1;
  } while (left < nums.length);

  return res;
};

console.log(72, maxProduct([ -1, 0, -1, 2, -3, 1, 2, 3, -2 ]));
