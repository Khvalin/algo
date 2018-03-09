/**
 * @param {number[]} A
 * @param {number} L
 * @param {number} R
 * @return {number}
 */
var numSubarrayBoundedMax = function(A, L, R) {
  let count = 0;

  let stackSize = 0;
  let max = Number.NEGATIVE_INFINITY;

  for (let i = 0; i < A.length; i++) {
    if ((A[i] >= L && A[i] <= R) || A[i] <= max) {
      if (max < A[i]) {
        max = A[i];
      }

      stackSize++;
      count += stackSize;

      if (A[i] < L) {
        count--;
      }
    } else {
      if (A[i] > max) {
        stackSize = 0;
        max = Number.NEGATIVE_INFINITY;
      }
    }

    console.log(A[i], count, stackSize);
  }

  return count;
};

console.log(numSubarrayBoundedMax([ 1, 2, 0, 0 ], 1, 2));

console.log(numSubarrayBoundedMax([ 73, 55, 36, 5, 55, 14, 9, 7, 72, 52 ], 32, 69)); // 22
