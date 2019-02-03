/**
 * Definition for an interval.
 * function Interval(start, end) {
 *     this.start = start;
 *     this.end = end;
 * }
 */
/**
 * @param {Interval[]} A
 * @param {Interval[]} B
 * @return {Interval[]}
 */
var intervalIntersection = function(A, B) {
  const res = []
  while (A.length && B.length) {
    let a = A[0]
    let b = B[0]

    let l = Math.max(a[0], b[0])
    let r = Math.min(a[1], b[1])
    if (l <= r) {
      if (res.length === 0 || res[res.length - 1][1] < l) {
        res.push([ l, r ])
      } else {
        res[res.length - 1][1] = r
      }
    }

    if (a[1] === b[1]) {
      A.shift()
      B.shift()
    } else {
      if (a[1] > b[1]) {
        B.shift()
      } else {
        A.shift()
      }
    }
  }

  return res
}

console.log(
  intervalIntersection(
    [ [ 0, 2 ], [ 5, 10 ], [ 13, 23 ], [ 24, 25 ] ], //
    [ [ 1, 5 ], [ 8, 12 ], [ 15, 24 ], [ 25, 26 ] ]
  )
)
