/**
 * 
 *Given an array of integers, return the maximum sum for a non-empty subarray (contiguous elements) with at most one element deletion. In other words, you want to choose a subarray and optionally delete one element from it so that there is still at least one element left and the sum of the remaining elements is maximum possible.

Note that the subarray needs to be non-empty after deleting one element.
 */

//Just maintain two variables, sum is the current maximal sum without any delete and delete is the current maximal sum with just one delete.

/**
 * @param {number[]} arr
 * @return {number}
 */
var maximumSum = function(arr) {
  let res = arr[0]
  let [ sum, d ] = [ 0, 0 ]

  for (const n of arr) {
    if (n < 0) {
      res = Math.max(res, n)
      d = Math.max(sum, d + n)
      sum = Math.max(sum + n, 0)
    } else {
      sum += n
      d += n
      res = Math.max(res, d, sum)
    }
  }

  return res
}

let res = maximumSum([ -5, -6, -23, 14, 1, -8, -1, 8, 1, -16, 15, 22, 21, 14, 5, -4, -3, 14, -1, 6, -17, 22, -19 ])
console.log(res)

res = maximumSum([ 10, -2, -2, 30 ])
console.log(res)

res = maximumSum([ -1, -1, -1, -1 ])
console.log(res)
