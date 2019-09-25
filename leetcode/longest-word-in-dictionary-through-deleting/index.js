/**
 * @param {string} s
 * @param {string[]} d
 * @return {string}
 */
var findLongestWord = function(s, d) {
  const next = {}

  let res = ''
  for (let i = 0; i < d.length; i++) {
    const w = d[i]
    const ch = w[0]
    next[ch] = next[ch] || []
    next[ch].push({ wi: i, ci: 0 })
  }

  for (const ch of s) {
    next[ch] = next[ch] || []
    const len = next[ch].length

    for (let i = 0; i < len; i++) {
      const p = next[ch].shift()
      p.ci++

      const w = d[p.wi]
      if (p.ci >= w.length) {
        if (res.length < w.length || (res.length === w.length && res.localeCompare(w) > 0)) {
          res = w
          continue
        }
      }

      const c = w[p.ci]
      next[c] = next[c] || []
      next[c].push(p)
    }
  }

  return res
}

console.log(findLongestWord('abpcplea', [ 'ale', 'apple', 'monkey', 'plea' ]))
console.log(findLongestWord('abpcplea', [ 'a', 'b', 'c' ]))
