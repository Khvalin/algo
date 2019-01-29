/**
 * @param {number} A
 * @param {number} B
 * @return {string}
 */
var strWithout3a3b = function(A, B) {
  let s = ''
  let bCount = 0
  let aCount = 0

  while (A + B > 0) {
    if (bCount == 2) {
      A--
      s += 'a'
      bCount = 0
      aCount = 1
      continue
    }

    if (aCount == 2) {
      B--
      s += 'b'
      aCount = 0
      bCount = 1
      continue
    }
    if (A > B) {
      s += 'a'
      bCount = 0
      aCount++
      A--
      continue
    }

    if (B > A) {
      s += 'b'
      aCount = 0
      bCount++
      B--
      continue
    }
  }

  return s
}
