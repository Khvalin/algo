/**
 * @param {string[]} words
 * @param {string} order
 * @return {boolean}
 */
var isAlienSorted = function(words, order) {
  const ABC = 'abcdefghijklmnopqrstuvwxyz'

  words = words.map((w) => {
    let res = ''
    for (let i = 0; i < w.length; i++) {
      res += ABC[order.indexOf(w[i])]
    }

    return res
  })

  const sortedWords = [ ...words ]
  sortedWords.sort()
  for (let i = 0; i < words.length; i++) {
    if (sortedWords[i] !== words[i]) {
      return false
    }
  }

  return true
}

console.log(isAlienSorted([ 'apple', 'app' ], 'abcdefghijklmnopqrstuvwxyz'))
console.log(isAlienSorted([ 'hello', 'leetcode' ], 'hlabcdefgijkmnopqrstuvwxyz'))
