/**
 * @param {string} S
 * @return {string}
 */
var removeOuterParentheses = function(S) {
  const parts = []
  let i = 0
  let next = ''
  for (const ch of S) {
    if (ch === '(') {
      i++
    }
    if (ch === ')') {
      i--
    }

    next += ch

    if (i == 0) {
      parts.push(next)
      next = ''
    }
  }

  next && parts.push(next)

  let res = ''

  for (let s of parts) {
    let n = 0
    for (let i = 0; i < s.length - 1; i++) {
      let ch = s[i]

      if (ch === '(') {
        n++
      }
      if (ch === ')') {
        n--
      }
    }

    if (n == 1) {
      s = s.substr(1, s.length - 2)
    }

    res += s
  }

  return res
}

console.log(removeOuterParentheses('(()())(())'))
console.log(removeOuterParentheses('(()())(())(()(()))'))
//console.log(removeOuterParentheses('(()())(())(()(()))'))
