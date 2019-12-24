/**
 * @param {number[]} nums
 * @param {number} k
 * @return {boolean}
 */
var isPossibleDivide = function(nums, k) {
  const d = {};

  // count occurances and store into a map.
  for (const n of nums) {
    d[n] = d[n] || 0;
    d[n]++;
  }

  // "root" is the starting number of each sequence.
  const roots = [];
  for (const n of nums) {
    if (!d[n - 1]) {
      roots.push(n);
    }
  }
  console.log(roots);

  // test each "root".
  while (roots.length > 0) {
    const r = roots.shift();
    if (d[r] <= 0) {
      continue;
    }

    // visit every number in [r..r+k] range and attempt building sequence, note we test k + 1 consecutive numbers to make sure we check all potential "roots".
    for (let i = r; i <= r + k; i++) {
      if (i < r + k) {
        if (!d[i]) {
          return false;
        }
        d[i]--;
      }

      // keeping track of new "roots".
      if (d[i] && !d[i - 1]) {
        roots.push(i);
      }
    }
  }

  return true;
};

let nums, k, r;

// nums = [3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11];
// k = 3;
// r = isPossibleDivide(nums, k);
// console.log(r);

nums = [ 4, 5, 3, 3, 3, 2, 2, 1, 1 ];
k = 3;
r = isPossibleDivide(nums, k);
console.log(r);

// nums = [15, 16, 17, 18, 19, 16, 17, 18, 19, 20, 6, 7, 8, 9, 10, 3, 4, 5, 6, 20];
// k = 5;
// r = isPossibleDivide(nums, k);
// console.log(r);

// nums = [10, 9, 8, 1, 2, 3, 2, 3, 4, 4, 5, 6, 10, 11, 12];
// k = 3;
// r = isPossibleDivide(nums, k);
// console.log(r);

// nums = [9, 10, 11, 5, 6, 7, 5, 6, 7, 7, 8, 9];
// k = 3;
// r = isPossibleDivide(nums, k);
// console.log(r);

// nums = [9, 8, 18, 20, 12, 13, 12, 20, 14, 22, 17, 19, 11, 14, 8, 25, 14, 11, 8, 17, 23, 26, 12, 6, 14, 9, 20, 4, 11, 3, 16, 10, 8, 10, 5, 22, 18, 21, 18, 12, 17, 17, 7, 19, 24, 15, 9, 22, 11, 10, 2, 18, 7, 23, 24, 11, 21, 15, 5, 18, 13, 15, 26, 2, 20, 8, 16, 10, 27, 7, 24, 6, 3, 23, 9, 17, 7, 9, 21, 20, 15, 21, 5, 13, 19, 16, 12, 16, 21, 16, 19, 19, 15, 25, 4, 14, 6, 13, 13, 10];
// k = 20;
// r = isPossibleDivide(nums, k);
// console.log(r);

// prettier-ignore
// nums = [9, 13, 15, 23, 22, 25, 4, 4, 29, 15, 8, 23, 12, 19, 24, 17, 18, 11, 22, 24, 17, 17, 10, 23, 21, 18, 14, 18, 7, 6, 3, 6, 19, 11, 16, 11, 12, 13, 8, 26, 17, 20, 13, 19, 22, 21, 27, 9, 20, 15, 20, 27, 8, 13, 25, 23, 22, 15, 9, 14, 20, 10, 6, 5, 14, 12, 7, 16, 21, 18, 21, 24, 23, 10, 21, 16, 18, 16, 18, 5, 20, 19, 20, 10, 14, 26, 2, 9, 19, 12, 28, 17, 5, 7, 25, 22, 16, 17, 21, 11]
// k = 10;
// r = isPossibleDivide(nums, k);
// console.log(r);
