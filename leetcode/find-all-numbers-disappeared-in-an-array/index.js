/**
 * @param {number[]} nums
 * @return {number[]}
 */
var findDisappearedNumbers = function(nums) {
  nums.unshift(0)

  for (let i = 1; i < nums.length;i++) {
    let j = nums[i]
    while (nums[j] > 0) {
      const t = nums[j]
      nums[j] = 0
      j = t
    }
  }

  const res = []
  for (let i = 1; i < nums.length;i++) {
    if (nums[i] > 0) {
      res.push(i)
    }
  }

  return res
};