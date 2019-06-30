/**
 * @param {number} candies
 * @param {number} num_people
 * @return {number[]}
 */
var distributeCandies = function(candies, num_people) {
  const res = []
  for (let i = 0; i < num_people; i++) {
    res.push(0)
  }
  let n = 1
  let i = 0
  while (candies > 0) {
    n = Math.min(n, candies)
    res[i % num_people] += n
    candies -= n
    i++
    n++
  }

  return res
}

console.log(distributeCandies(7, 4))
console.log(distributeCandies(10, 3))
