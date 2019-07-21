/**
 * @param {number} N
 * @param {number[][]} trust
 * @return {number}
 */
var findJudge = function(N, trust) {
  let T1 = {}
  let T0 = {}
  let indices = new Set()
  for (const t of trust) {
    indices.add(t[0])
    indices.add(t[1])

    T0[t[0]] = T0[t[0]] || 0
    T1[t[1]] = T1[t[1]] || 0

    T1[t[1]]++
    T0[t[0]]++
  }

  /*   console.log(indices.values())
  console.log(T0, T1) */

  let r = N === 1 && trust.length === 0 ? 1 : -1
  for (let n of indices.values()) {
    if (T1[n] === N - 1 && !T0[n]) {
      if (r == -1) {
        r = n
      } else {
        r = -1
        break
      }
    }
  }

  return r
}

console.log(findJudge(2, [ [ 1, 2 ], [ 2, 1 ] ]))
console.log(findJudge(4, [ [ 1, 3 ], [ 1, 4 ], [ 2, 3 ], [ 2, 4 ], [ 4, 3 ] ]))
console.log(findJudge(1, []))
