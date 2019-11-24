/**
 * @param {number[][]} points
 * @return {number}
 */
var minTimeToVisitAllPoints = function(points) {
  let r = 0

  let a = points[0]
  for (let i = 1; i < points.length; i++) {
    b = points[i]

    let [ xa, ya ] = a
    let [ xb, yb ] = b

    const d = Math.min(Math.abs(xb - xa), Math.abs(yb - ya))

    xa += d * Math.sign(xb - xa)
    ya += d * Math.sign(yb - ya)

    r += Math.abs(d) + Math.abs(xb - xa) + Math.abs(yb - ya)

    a = [ ...b ]
  }

  return r
}

let points = [ [ 1, 1 ], [ 3, 4 ], [ -1, 0 ] ]
console.log(7, minTimeToVisitAllPoints(points))

points = [ [ 3, 2 ], [ -2, 2 ] ]
console.log(5, minTimeToVisitAllPoints(points))
