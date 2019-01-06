/**
 * @param {number} x
 * @param {number} y
 * @param {number} bound
 * @return {number[]}
 */
var powerfulIntegers = function(x, y, bound) {
  //
  let res = new Set()
  let i = 0

  let first = 1
  while (first < bound) {
    let sum = first
    let j = 0
    while (sum < bound) {
      sum = first + Math.pow(y, j)
      if (sum <= bound) {
        res.add(sum)
      }
      j++
      if (y == 1) {
        break
      }
    }

    first = Math.pow(x, i)
    if (x == 1) {
      break
    }
    i++
  }

  return [ ...res.values() ]
}

console.log(powerfulIntegers(3, 5, 15))
