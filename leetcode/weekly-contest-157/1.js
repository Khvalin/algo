/**
 * @param {number[]} chips
 * @return {number}
 */
var minCostToMoveChips = function(chips) {
  let [ odd, even ] = [ 0, 0 ]

  for (const n of chips) {
    odd += n % 2
    even += !(n % 2)
  }

  return Math.min(odd, even)
}

let input
input = [ 1, 2, 3 ]
console.log(minCostToMoveChips(input))

console.log(minCostToMoveChips([ 2, 2, 2, 3, 3 ]))

// prettier-ignore
console.log(minCostToMoveChips([1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30]))
