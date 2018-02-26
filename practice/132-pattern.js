/**
 * @param {number[]} nums
 * @return {boolean}
 */
var find132pattern = function(nums) {
  let i = 0;
  let j = 1;

  //i < j < k and ai < ak < aj.
  while (j < nums.length - 1) {
    if (nums[j] > nums[i]) {
      while (nums[j + 1] > nums[j] && j < nums.length - 1) {
        j++;
      }

      for (let k = j + 1; k < nums.length; k++) {
        if (nums[i] < nums[k] && nums[k] < nums[j]) {
          return true;
        }
      }
    }
    if (nums[i] > nums[j]) {
      i = j;
    }
    j++;
  }

  return false;
};

console.log(find132pattern([ 3, 5, 0, 3, 4 ]));
console.log(find132pattern([ 3, 5, 0, 2, 1 ]));
//console.log(find132pattern([ 1, 0, 1, -4, -3 ]));
//console.log(find132pattern([ 3, 1, 4, 2 ]));
//console.log(find132pattern([ -2, -1, 0, 1, 0, 1, -1 ]));
