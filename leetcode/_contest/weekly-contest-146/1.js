/**
 * @param {number[][]} dominoes
 * @return {number}
 */
var numEquivDominoPairs = function(dominoes) {
  let res = 0
  let h = {}

  for (let d of dominoes) {
    d = d.sort((a, b) => a - b)
    const s = d[0] * 10 + d[1]
    h[s] = h[s] || 0
    if (h[s]) {
      res += h[s]
    }
    h[s]++
  }

  return res
}

console.log(numEquivDominoPairs([ [ 1, 2 ], [ 1, 2 ], [ 1, 1 ], [ 1, 2 ], [ 2, 2 ] ]))
console.log(numEquivDominoPairs([ [ 1, 2 ], [ 2, 1 ], [ 3, 4 ], [ 5, 6 ] ]))
