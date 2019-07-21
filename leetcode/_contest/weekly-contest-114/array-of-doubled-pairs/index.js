/**
 * @param {number[]} A
 * @return {boolean}
 */
var canReorderDoubled = function(A) {
  A.sort((a, b) => a - b) //[-4,-2,2,4]
  const M = {}
  for (const n of A) {
    if (!(n in M)) {
      M[n] = 0
    }
    M[n]++
  }

  for (const n of A) {
    if (M[n] > 0) {
      M[n]--
      let p = n < 0 ? n / 2 : n * 2

      if (M[p]) {
        M[p]--
      } else {
        return false
      }
    }
  }

  return true
}

console.log(canReorderDoubled([ 4, -2, 2, -4 ]))
console.log(canReorderDoubled([ 2, 1, 2, 6 ]))
console.log(canReorderDoubled([ 1, 2, 4, 16, 8, 4 ]))
