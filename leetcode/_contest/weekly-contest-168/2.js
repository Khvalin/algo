/**
 * @param {number[]} nums
 * @param {number} k
 * @return {boolean}
 */
var isPossibleDivide = function(nums, k) {
  if (nums.length % k > 0) {
    return false;
  }

  nums = nums.sort((a, b) => a - b);

  const l = nums.length / k;

  let a = [];
  for (let i = 0; i < l; i++) {
    a.push([]);
  }

  a[0].push(nums[0]);

  for (let i = 1; i < nums.length; i++) {
    let j = 0;
    while (j < a.length && (a[j].length > 0 && a[j][a[j].length - 1] !== nums[i] - 1)) {
      j++;
    }

    if (j >= a.length) {
      return false;
    }

    a[j].push(nums[i]);

    if (a[j].length === k) {
      a.splice(j, 1);
    }
  }

  return true;
};

let nums = [ 1, 2, 3, 3, 4, 4, 5, 6 ],
  k = 4;
let r = isPossibleDivide(nums, k);
console.log(r);

(nums = [ 3, 3, 2, 2, 1, 1 ]), (k = 3);
r = isPossibleDivide(nums, k);
console.log(r);

nums = [ 15, 16, 17, 18, 19, 16, 17, 18, 19, 20, 6, 7, 8, 9, 10, 3, 4, 5, 6, 20 ];
k = 5;
r = isPossibleDivide(nums, k);
console.log(r);
