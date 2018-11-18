/**
 * @param {string} S
 * @return {number[]}
 */
var diStringMatch = function(S) {
  let [ i, d ] = [ 0, S.length ]

  res = []
  for (const ch of S) {
    if (ch == 'I') {
      res.push(i)
      i++
    } else {
      res.push(d)
      d--
    }
  }
  res.push(d)

  return res
}

console.log(diStringMatch('IDID'))
console.log(diStringMatch('DDI'))
