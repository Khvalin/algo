/**
 * @param {number[][]} slots1
 * @param {number[][]} slots2
 * @param {number} duration
 * @return {number[]}
 */
var minAvailableDuration = function(slots1, slots2, duration) {
  const sortfn = (s1, s2) => {
    let r = s1[0] - s2[0]
    if (r !== 0) {
      return r
    }

    return s2[1] - s1[1]
  }
  slots1 = slots1.sort(sortfn)
  slots2 = slots2.sort(sortfn)

  /*  console.log(slots1)
  console.log(slots2) */

  let i = 0
  let j = 0

  const res = []
  while (res.length === 0 && i < slots1.length && j < slots2.length) {
    const [ s1, e1 ] = slots1[i]
    const [ s2, e2 ] = slots2[j]

    const l = Math.max(s1, s2)
    const r = Math.min(e1, e2)

    if (l + duration <= r) {
      res.push([ l, l + duration ])
    }

    if (slots1[i][1] < slots2[j][1]) {
      i++
    } else {
      j++
    }
  }

  return res
}

let slots1 = [ [ 10, 50 ], [ 60, 120 ], [ 140, 210 ] ],
  slots2 = [ [ 0, 15 ], [ 60, 70 ] ],
  duration = 8
console.log(minAvailableDuration([ [ 10, 50 ], [ 60, 120 ], [ 140, 210 ] ], [ [ 0, 15 ], [ 60, 70 ] ], 8))
/*
;(slots1 = [ [ 10, 50 ], [ 60, 120 ], [ 140, 210 ] ]), (slots2 = [ [ 0, 15 ], [ 60, 70 ] ]), (duration = 12)
console.log(minAvailableDuration(slots1, slots2, duration))
*/
