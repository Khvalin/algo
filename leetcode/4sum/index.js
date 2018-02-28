/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[][]}
 */
var fourSum = function(nums, target) {
  const result = new Map();
  const mask = [ 0, 1, 2, 3 ];
  while (mask[0] <= nums.length - 4) {
    const threeSum = nums[mask[0]] + nums[mask[1]] + nums[mask[2]];
    for (let i = mask[3]; i < nums.length; i++) {
      if (threeSum + nums[i] === target) {
        const val = [ nums[mask[0]], nums[mask[1]], nums[mask[2]], nums[i] ].sort();
        result.set(val.join(','), val);
      }
    }

    let d = 2;
    while (d > 0 && mask[d] >= nums.length - 4 + d) {
      d--;
    }

    mask[d]++;

    for (let j = d + 1; j < mask.length; j++) {
      mask[j] = mask[d] + j - d;
    }
    // console.log(mask);
    // break;
  }

  return Array.from(result.values());
};

console.log(fourSum([ 1, 0, -1, 0, -2, 2 ], 0));
console.log(fourSum([ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1 ], 4));
