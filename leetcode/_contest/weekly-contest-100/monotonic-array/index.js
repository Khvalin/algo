/**
 * @param {number[]} A
 * @return {boolean}
 */
var isMonotonic = function(A) {
  if (A.length === 1) {
    return true
  }

  let gSign = 0
  for (let i = 1; i < A.length; i++) {
    let sign = Math.sign(A[i] - A[i - 1])
    if (gSign == 0) {
      gSign = sign
    } else {
      if (sign != 0 && gSign != sign) {
        return false
      }
    }
  }

  return true
}
