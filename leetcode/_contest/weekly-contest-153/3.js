/**
 * @param {number[]} arr
 * @return {number}
 */
var maximumSum = function(arr) {
  let res = 0
  let d = null

  let r = Math.max(...arr)
  if (r < 0) {
    return r
  }

  let s = 0
  let sub = 0
  for (const n of arr) {
    if (n >= 0) {
      s += n
      sub += n
    } else {
      if (n < 0 && d === null) {
        if (s > 0) {
          d = n
        }
        sub = 0
      } else {
        if (d > n) {
          s += d
          d = n
        } else {
          s += n
        }

        if (sub > s) {
          s = sub
          d = n
        }
        sub = 0
      }
    }

    res = Math.max(res, s)
    console.log(n, s, d)
  }

  return res
}

let res = maximumSum([ -5, -6, -23, 14, 1, -8, -1, 8, 1, -16, 15, 22, 21, 14, 5, -4, -3, 14, -1, 6, -17, 22, -19 ])
console.log(res)
/* 
res = maximumSum([ 10, -2, -2, 30 ])
console.log(res)

res = maximumSum([ -1, -1, -1, -1 ])
console.log(res)
 */
