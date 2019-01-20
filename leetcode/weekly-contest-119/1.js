/**
 * @param {number[][]} points
 * @param {number} K
 * @return {number[][]}
 */
var kClosest = function(points, K) {
  const d = (p) => {
    return p[0] * p[0] + p[1] * p[1]
  }

  const a = []
  for (const p of points) {
    const pd = d(p)

    a.push({ d: pd, point: [ ...p ] })
  }

  a.sort((a, b) => a.d - b.d)

  const res = []
  for (let i = 0; i < K; i++) {
    res.push(a[i].point)
  }

  return res
}

console.log(kClosest([ [ 3, 3 ], [ 5, -1 ], [ -2, 4 ] ], 2))
