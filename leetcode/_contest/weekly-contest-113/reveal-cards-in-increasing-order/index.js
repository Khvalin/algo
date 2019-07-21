/**
 * @param {number[]} deck
 * @return {number[]}
 */
var deckRevealedIncreasing = function(deck) {
  let indices = deck.map((_, i) => i)
  const newIndices = []
  while (indices.length > 0) {
    newIndices.push(indices.shift())
    if (indices.length) {
      indices.push(indices.shift())
    }
  }
  //console.log(indices)
  //console.log(newIndices)

  deck.sort((a, b) => a - b)

  const r = Array.from(deck)
  while (deck.length) {
    r[newIndices.shift()] = deck.shift()
  }
  return r
}

console.log(deckRevealedIncreasing([ 17, 13, 11, 2, 3, 5, 7 ]))
console.log(deckRevealedIncreasing([ 17, 13, 11, 3, 5, 7 ]))
