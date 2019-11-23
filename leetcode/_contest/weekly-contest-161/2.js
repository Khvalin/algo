/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var numberOfSubarrays = function(nums, k) {
  let c = 0
  let r = 0

  let j = 0
  for (let i = 0; i < nums.length; i++) {
    const n = nums[i]

    if (n % 2 > 0) {
      c++
    }

    if (c === k) {
      j = i
      r++
    }
  }
}

console.log(numberOfSubarrays([ 1, 1, 2, 1, 1 ], 3))
