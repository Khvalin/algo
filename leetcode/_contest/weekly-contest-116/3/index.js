/**
 * @param {number[][]} points
 * @return {number}
 */
var minAreaFreeRect = function(points) {
  function round(value, precision = 7) {
    var multiplier = Math.pow(10, precision || 0)
    return Math.round(value * multiplier) / multiplier
  }

  const getDist = (a, b) => {
    return Math.sqrt(Math.pow(a[0] - b[0], 2) + Math.pow(a[1] - b[1], 2))
  }

  const distancies = []
  const distMap = {}
  const counts = {}

  for (let i = 0; i < points.length; i++) {
    for (let j = i + 1; j < points.length; j++) {
      const d = round(getDist(points[i], points[j]))
      counts[d] = counts[d] || 0
      counts[d]++

      distancies.push({ i1: i, i2: j, dist: d })
      distMap[i + ' ' + j] = d
    }
  }

  let res = null
  const cans = distancies.filter((d) => counts[d.dist] > 4)
  for (i = 0; i <= cans.length - 4; i++) {
    const col = [ cans[i], cans[i + 1], cans[i + 2], cans[i + 3] ]
    const a = col.map((c) => c.i1)
    const b = col.map((c) => c.i2)

    if (a[0] == a[1] && b[0] == a[2] && b[1] == b[3] && b[2] == a[3]) {
      const diag1 = distMap[a[0] + ' ' + b[2]]
      const diag2 = distMap[a[2] + ' ' + b[3]]

      if (diag1 == diag2) {
        const sq = col[0].dist * col[3].dist
        res = res || sq
        res = Math.min(res, sq)
      }
    }
  }

  return res || 0

  //console.log(distancies)
  //console.log(candidates)
}

console.log(minAreaFreeRect([ [ 0, 1 ], [ 2, 1 ], [ 1, 1 ], [ 1, 0 ], [ 2, 0 ] ]))
console.log(minAreaFreeRect([ [ 0, 1 ], [ 2, 1 ], [ 1, 1 ], [ 1, 0 ], [ 2, 0 ] ]))
console.log(minAreaFreeRect([ [ 3, 1 ], [ 1, 1 ], [ 0, 1 ], [ 2, 1 ], [ 3, 3 ], [ 3, 2 ], [ 0, 2 ], [ 2, 3 ] ]))
