/**
 * @param {number[]} arr
 * @return {void} Do not return anything, modify arr in-place instead.
 */
var duplicateZeros = function(arr) {
  const cache = []
  for (let i = 0; i < arr.length; i++) {
    const n = arr[i]

    cache.push(n)
    if (n == 0) {
      cache.push(0)
    }

    if (cache.length > 0) {
      arr[i] = cache.shift()
    }
  }
}

var a
a = [ 1, 0, 2, 3, 0, 4, 5, 0 ]
duplicateZeros(a)
console.log(a)

a = [ 1, 2, 3 ]
duplicateZeros(a)
console.log(a)

a = [ 0, 1, 7, 6, 0, 2, 0, 7 ]
duplicateZeros(a)
console.log(a)
