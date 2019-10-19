/**
 * @param {number[]} sweetness
 * @param {number} K
 * @return {number}
 */
var maximizeSweetness = function(sweetness, K) {
  const L = sweetness.length
  const memo = {}

  const calc = (ind, min, k, cuts = []) => {
    const key = `${ind} ${min} ${k}`

    if (key in memo) {
      return memo[key]
    }

    if (k > 0 && ind >= L) {
      return Number.NEGATIVE_INFINITY
    }

    if (k <= 0) {
      return Number.NEGATIVE_INFINITY
    }

    if (k === 1) {
      let s = 0
      for (let i = ind; i < L; i++) {
        s += sweetness[i]
      }

      const r = Math.min(min, s)

      memo[key] = r
      return r
    }

    let s = 0
    let r = Number.NEGATIVE_INFINITY
    for (let i = ind; i < L; i++) {
      s += sweetness[i]
      if (s < min) {
        min = s
      }
      const t = calc(i + 1, min, k - 1, [ ...cuts, s ])

      r = Math.max(r, t)
      console.log(s)
    }

    //memo[key] = r

    return r
  }

  const r = calc(0, Number.POSITIVE_INFINITY, K)

  return r
}

console.log(maximizeSweetness([ 1, 2, 3, 4, 5, 6, 7, 8, 9 ], 5))
