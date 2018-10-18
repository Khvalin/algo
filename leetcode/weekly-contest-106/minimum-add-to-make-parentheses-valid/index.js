/**
 * @param {string} S
 * @return {number}
 */
var minAddToMakeValid = function(S) {
  let b = 0
  let res = 0
  for (ch of S) {
    if (ch == '(') {
      b++
    } else {
      b--
    }

    if (b < 0) {
      res++
      b = 0
    }
  }

  if (b > 0) {
    res += b
  }

  return res
}
