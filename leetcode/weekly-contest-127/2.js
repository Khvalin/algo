'use strict'

/**
 * @param {number} N
 * @return {number}
 */
var clumsy = function(N) {
  const ops = [ '*', '/', '+', '-' ]
  let a = []
  let j = 0
  for (let i = N; i > 0; i--) {
    if (a[a.length - 1] == '*') {
      a.pop()
      a[a.length - 1] *= i
    } else {
      if (a[a.length - 1] == '/') {
        a.pop()
        a[a.length - 1] = Math.floor(a[a.length - 1] / i)
      } else {
        a.push(i.toString())
      }
    }

    if (i > 1) {
      a.push(ops[j])
      j = (j + 1) % 4
    }
  }
  console.log(a)
  return eval(a.join(''))

  //console.log(a)
}

console.log(clumsy(10))
