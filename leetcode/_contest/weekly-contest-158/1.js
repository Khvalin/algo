/**
 * @param {string} s
 * @return {number}
 */
var balancedStringSplit = function(s) {
  let [ l, r ] = [ 0, 0 ]
  let res = 0

  for (const ch of s) {
    if (ch === 'L') {
      l++
    } else {
      r++
    }

    if (l === r) {
      res++
    }
  }

  return res
}

console.log(balancedStringSplit('RLRRLLRLRL'))
