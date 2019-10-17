/**
 * @param {number} n
 * @param {number[]} rollMax
 * @return {number}
 */
var dieSimulator = function(n, rollMax) {
  //Input: n = 2, rollMax = [1,1,2,2,2,3]
  /*
    12
    21
    33
  */

  const NMAX = 1000000007

  let roll = [ 1, 1, 1, 1, 1, 1 ]

  for (let i = 1; i < n; i++) {
    const prev = [ ...roll ]
    for (let j = 0; j < roll.length; j++) {
      let s = 0
      for (let k = 0; k < 6; k++) {
        s += prev[k]
        if (i % rollMax[k] === 0 && j === k) {
          s -= 1
        }
      }

      roll[j] = s % NMAX
    }

    console.log(roll)
  }

  let r = 0
  for (const n of roll) {
    r += n
    r %= NMAX
  }

  return r
}

let n = 2,
  rollMax = [ 1, 1, 1, 2, 2, 3 ]
console.log(dieSimulator(n, rollMax))

console.log(dieSimulator(2, [ 1, 1, 1, 1, 1, 1 ]))

console.log(dieSimulator(3, [ 1, 1, 1, 20, 2, 3 ]))
