/**
 * @param {number[][]} costs
 * @return {number}
 */
var twoCitySchedCost = function(costs) {
  const N = costs.length
  let diff = []

  let r = 0
  for (const cost of costs) {
    r += cost[0]
    diff.push(cost[0] - cost[1])
  }

  diff = diff.sort((a, b) => a - b)
  for (let i = N / 2; i < N; i++) {
    r -= diff[i]
  }

  return r
}

let r
r = twoCitySchedCost([ [ 10, 20 ], [ 30, 200 ], [ 400, 50 ], [ 30, 20 ] ])
console.log(r)

r = twoCitySchedCost([ [ 259, 770 ], [ 448, 54 ], [ 926, 667 ], [ 184, 139 ], [ 840, 118 ], [ 577, 469 ] ])
console.log(r)
