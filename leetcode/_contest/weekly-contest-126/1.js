/**
 * @param {string[]} A
 * @return {string[]}
 */
var commonChars = function(A) {
  let res = []

  const countChars = (s) => {
    const res = {}
    for (let i = 0; i < s.length; i++) {
      let ch = s[i]
      res[ch] = res[ch] || 0
      res[ch]++
    }

    return res
  }

  let chars = countChars(A[0])
  for (let i = 1; i < A.length; i++) {
    let map = countChars(A[i])
    const newChars = {}

    for (const ch in map) {
      if (ch in chars) {
        newChars[ch] = Math.min(chars[ch], map[ch])
      }
    }

    chars = Object.assign({}, newChars)
  }

  for (const ch in chars) {
    for (let i = 0; i < chars[ch]; i++) {
      res.push(ch)
    }
  }

  return res
}

console.log(commonChars([ 'bella', 'label', 'roller' ]))
console.log(commonChars([ 'cool', 'lock', 'cook' ]))
