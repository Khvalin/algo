/**
 * @param {string} s
 * @return {number}
 */
var balancedString = function(s) {
  const counts = {}
  const l = s.length / 4

  const overflow = {}
  const d = {}

  for (let i = 0; i < s.length; i++) {
    const ch = s[i]

    counts[ch] = counts[ch] || 0
    d[ch] = d[ch] || [ -1 ]
    counts[ch]++

    d[ch][counts[ch]] = i

    if (counts[ch] > l) {
      overflow[ch] = counts[ch] - l
    }
  }

  for (const ch of 'QWER') {
    counts[ch] = 0
  }

  if (Object.keys(overflow).length === 0) {
    return 0
  }

  let r = Number.POSITIVE_INFINITY

  outer: for (let i = 0; i < s.length; i++) {
    const ch = s[i]

    counts[ch]++

    if (!(ch in overflow)) {
      continue
    }

    let len = 1
    for (const c in overflow) {
      if (counts[c] < overflow[c]) {
        continue outer
      }

      //  console.log(i, counts[c], overflow[c], d[c][counts[c] - overflow[c]])

      if (c !== ch || overflow[c] > 1) {
        len = Math.max(len, i - d[c][counts[c] - overflow[c]])
      }
    }

    r = Math.min(len, r)
  }

  return r
}

console.log(3, balancedString('WQWRQQQW'))
console.log(4, balancedString('WWQQRRRRQRQQ'))

console.log(balancedString('QEQEWWEQ'))
console.log(balancedString('QWER'))
console.log(balancedString('QQWE'))

console.log(balancedString('QQQQ'))
