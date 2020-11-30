class Solution {
  solve(nums) {
    let n = 0;
    for (let i = 0; i < nums.length; i++) {
      if (n >= i) {
        n = Math.max(n, i + nums[i]);
      }
    }

    return n >= nums.length - 1;
  }
}
