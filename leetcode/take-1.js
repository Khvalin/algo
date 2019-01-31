/**
 * @param {number[]} nums
 * @return {number}
 */
var maxSubArray = function(nums) {
  let a = []

  let res = Math.max(...nums)
  for(let i = 0; i < nums.length;i++) {
    let s = nums[i]
    while (i < nums.length -1 && (nums[i+1] === 0 ||  Math.sign(nums[i]) === Math.sign(nums[i+1]))) {
      i++
      s += nums[i]
    }

    a.push(s)
  }

  let sum = 0

  for(let i = 0; i < a.length;i++) {
    if (sum + a[i] > 0) {
      sum += a[i]
      res = Math.max(res, sum)
    } else {
      res = Math.max(res, a[i])
      sum = 0
    }
  }

  return res
};    