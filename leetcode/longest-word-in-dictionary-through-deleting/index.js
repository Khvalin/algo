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
    const cur = [ ...next[ch] ]
    next[ch] = []

    for (const p of cur) {
      p.ci++

      const w = d[p.wi]
      if (p.ci >= w.length) {
        if (res.length < w.length || (res.length === w.length && res.localeCompare(w) > 0)) {
          res = w
          continue
        }
      }

      const ch = w[p.ci]
      next[ch] = next[ch] || []
      next[ch].push(p)
    }
  }

  return res
}

console.log(findLongestWord('abpcplea', [ 'ale', 'apple', 'monkey', 'plea' ]))
console.log(findLongestWord('abpcplea', [ 'a', 'b', 'c' ]))
