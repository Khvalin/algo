/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
var merge = function (nums1, m, nums2, n) {
  let i = m - 1;
  let j = n - 1;
  let k = m + n - 1;
  const min = Math.min(nums1[0] || 0, nums2[0] || 0) - 1;

  while (k >= 0) {
    let a = min;
    let b = min;
    if (i >= 0) {
      a = nums1[i];
    }

    if (j >= 0) {
      b = nums2[j];
    }

    if (a >= b) {
      i--;
    } else {
      j--;
    }

    nums1[k] = Math.max(a, b);
    k--;
  }
};
