/**
 * @param {string} ex
 * @return {boolean}
 */
var parseBoolExpr = function(ex) {
  if (ex.startsWith('!')) {
    ex = ex.substring(2, ex.length - 1)
    return !parseBoolExpr(ex)
  }

  if (ex.startsWith('(')) {
    ex = ex.substring(1, ex.length - 1)
  }

  if (ex === 'f') {
    return false
  }

  if (ex === 't') {
    return true
  }

  if (ex[0] === '|' || ex[0] === '&') {
    let res = ex[0] === '&'
    let sum = 0
    let prev = 2

    for (let i = 2; i < ex.length; i++) {
      const ch = ex[i]
      if (ch === '(') {
        sum--
        continue
      }
      if (ch === ')') {
        sum++
      }
      //console.log(ch, sum)

      if ((ch === ',' && sum === 0) || (ch === ')' && sum === 1)) {
        const expr = ex.substring(prev, i)
        let p = parseBoolExpr(expr)

        if (ex[0] === '&') {
          res = res && p
        } else {
          res = res || p
        }

        prev = i + 1
      }
    }
    return res
  }

  return !!(ex.length % 2)
}

/* console.log(parseBoolExpr('!(f)'))
console.log(parseBoolExpr('|(&(t,f,t),!(t))'))
console.log(parseBoolExpr('&(t,f)'))
 */

let r = parseBoolExpr(
  '!(!(!(!(&(&(&(&(f),&(!(t),&(f),|(f)),&(!(&(f)),&(t),|(f,f,t))),|(t),&(!(t),!(|(f,f,t)),!(&(f)))),!(&(|(f,f,t),&(&(f),&(!(t),&(f),|(f)),&(!(&(f)),&(t),|(f,f,t))),&(t))),&(&(&(!(&(f)),|(t),&(!(t),!(|(f,f,t)),!(&(f)))),!(|(f,f,t)),&(t,t,f)),&(f),&(&(t),&(!(t),!(|(f,f,t)),!(&(f))),|(f,f,t))))))))'
)

console.log(false, r)
