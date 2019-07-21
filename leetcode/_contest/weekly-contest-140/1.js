/**
 * @param {string} text
 * @param {string} first
 * @param {string} second
 * @return {string[]}
 */
var findOcurrences = function(text, first, second) {
  const words = text.match(/\b(\w+)\b/g)
  const res = []
  for (let i = 2; i < words.length; i++) {
    if (words[i - 1] === second && words[i - 2] === first) {
      res.push(words[i])
    }
  }

  return res
}

console.log(findOcurrences('alice is a good girl she is a good student', 'a', 'good'))
console.log(findOcurrences('we will we will rock you', 'we', 'will'))
