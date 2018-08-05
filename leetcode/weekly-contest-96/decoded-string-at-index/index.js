/**
 * @param {string} S
 * @param {number} K
 * @return {string}
 */
var decodeAtIndex = function(S, K) {
  const num = S.match(/\d+/g) || []
  const words = S.match(/[a-zA-Z]+/g)

  if (num.length == 0) {
    return S[K - 1]
  }

  let i = 0
  let len = 0
  while (i < words.length && K > (len = words[i].length * num[i])) {
    i++
    K -= len
  }

  if (words[i]) {
    K %= words[i].length
    const w = words[i] + words[i]
    return w[K - 1 + words[i].length]
  }

  return ''
}

console.log(decodeAtIndex('leet2code3', 10))
console.log(decodeAtIndex('leet2code3', 8))
console.log(decodeAtIndex('ha22', 5))
console.log(decodeAtIndex('a2345678999999999999999', 1))
console.log(decodeAtIndex('a2b3c4d5e6f7g8h9', 69280))
