/**
 * @param {number[]} arr
 * @return {number}
 */
var missingNumber = function(arr) {
  const d = (arr[arr.length - 1] - arr[0]) / arr.length

  for (let i = 1; i < arr.length; i++) {
    if (arr[i] !== arr[i - 1] + d) {
      return arr[i - 1] + d
    }
  }

  return a[0]
}

console.log(missingNumber([ 5, 7, 11, 13 ]))
console.log(missingNumber([ 15, 13, 12 ]))
