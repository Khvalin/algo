/**
 * @param {number[]} customers
 * @param {number[]} grumpy
 * @param {number} X
 * @return {number}
 */
var maxSatisfied = function(customers, grumpy, X) {
  let sub = 0

  let [ r, gSum ] = [ 0, 0 ]
  let j = 0

  for (let i = 0; i < customers.length; i++) {
    if (grumpy[i]) {
      gSum += customers[i]
      for (; j <= i - X; j++) {
        gSum -= customers[j] * grumpy[j]
      }
    } else {
      sub += customers[i]
    }

    r = Math.max(r, gSum)
  }

  return sub + r
}

let ans = maxSatisfied([ 1, 0, 1, 2, 1, 1, 7, 5 ], [ 0, 1, 0, 1, 0, 1, 0, 1 ], 3)
console.log(ans)

ans = maxSatisfied([ 4, 10, 10 ], [ 1, 1, 0 ], 2)
console.log(ans)
