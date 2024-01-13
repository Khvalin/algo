/**
 * @param {number[]} nums
 * @param {number} lower
 * @param {number} upper
 * @return {number}
 */
var countFairPairs = function (nums, lower, upper) {
  const smaller = (val) => {
    let l = 0, r = nums.length - 1;
    let res = 0;
    while (l < r) {
      const sum = nums[l] + nums[r];
      if (sum <= val) {
        res += r - l;
        l++;
        continue;
      }

      // if sum is greater, decrement right pointer
      r--;
    }

    return res;
  };

  nums.sort((a, b) => a - b);

  return smaller(upper) - smaller(lower - 1);
};

console.log(countFairPairs([1, 7, 9, 2, 5], 11, 11));
console.log(countFairPairs([0, 0, 0, 0, 0, 0], 0, 1));
