/**
   * @param {string} A
   * @param {string} B
   * @return {string[]}
   */
var uncommonFromSentences = function(A, B) {
  const aWords = A.match(/\S+/g)
  const bWords = B.match(/\S+/g)

  const aMap = {}
  for (word of aWords) {
    if (!aMap[word]) {
      aMap[word] = 0
    }
    aMap[word]++
  }

  const bMap = {}
  for (word of bWords) {
    if (!bMap[word]) {
      bMap[word] = 0
    }
    bMap[word]++
  }

  const result = []
  for (word of [ ...aWords, ...bWords ]) {
    if (aMap[word] === 1 && !bMap[word]) {
      result.push(word)
    }

    if (bMap[word] === 1 && !aMap[word]) {
      result.push(word)
    }
  }

  return result
}

console.log(uncommonFromSentences('apple apple', 'banana'))
console.log(uncommonFromSentences('this apple is sweet', 'this apple is sour'))
