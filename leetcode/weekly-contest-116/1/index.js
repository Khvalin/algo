/**
 * @param {number[]} A
 * @return {number}
 */
var repeatedNTimes = function(A) {
  const count = {}
  for (const n of A) {
    count[n] = count[n] || 0
    count[n]++
  }

  for (const n of A) {
    if (count[n] === A.length / 2) {
      return n
    }
  }

  return null
}

console.log(repeatedNTimes([ 1, 2, 3, 3 ]))
