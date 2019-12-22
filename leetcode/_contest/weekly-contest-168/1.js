/**
 * @param {number[]} nums
 * @return {number}
 */
var findNumbers = function(nums) {
  let res = 0;
  for (const n of nums) {
    if (Math.floor(Math.log10(n)) % 2 > 0) {
      res++;
    }
  }

  return res;
};

let r = findNumbers([ 12, 345, 2, 6, 7896 ]);
console.log(r);
