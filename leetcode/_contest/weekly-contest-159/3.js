/**
 * @param {string} s
 * @return {number}
 */
var balancedString = function(s) {
  const counts = {}
  const l = s.length / 4

  for (let i = 0; i < s.length; i++) {
    const ch = s[i]

    counts[ch] = counts[ch] || 0
    counts[ch]++
  }

  let r = Number.POSITIVE_INFINITY
  let j = 0

  for (let i = 0; i < s.length; i++) {
    const ch = s[i]

    counts[ch]--

    while (j < s.length && Object.values(counts).every((c) => c <= l)) {
      r = Math.min(r, i - j + 1)
      counts[s[j++]]++
    }
  }

  return r
}

console.log(balancedString('QQQQ'))

console.log(4, balancedString('WWQQRRRRQRQQ'))

console.log(3, balancedString('WQWRQQQW'))

console.log(balancedString('QEQEWWEQ'))
console.log(balancedString('QWER'))
console.log(balancedString('QQWE'))
