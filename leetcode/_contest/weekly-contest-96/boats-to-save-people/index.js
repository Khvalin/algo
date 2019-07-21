/**
 * @param {number[]} people
 * @param {number} limit
 * @return {number}
 */
var numRescueBoats = function(people, limit) {
  people = people.sort((a, b) => a - b)
  let result = 0
  while (people.length > 1) {
    // console.log(people)
    const p = people.pop()
    if (p + people[0] <= limit) {
      people.shift()
    }
    result++
  }

  result += people.length

  return result
}

console.log(numRescueBoats([ 1, 2 ], 3))
console.log(numRescueBoats([ 3, 5, 3, 4 ], 5))
console.log(numRescueBoats([ 3, 2, 2, 1 ], 3))
