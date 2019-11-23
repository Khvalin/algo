/**
 * @param {string} s1
 * @param {string} s2
 * @return {number}
 */
var minimumSwap = function(s1, s2) {
  if (s1.length !== s2.length) {
    return -1
  }

  let r1 = 0
  let r2 = 0
  let [ x1, x2, y1, y2 ] = [ 0, 0, 0, 0 ]
  for (let i = 0; i < s1.length; i++) {
    if (s1[i] === 'x') {
      x1++
    } else {
      y1++
    }

    if (s2[i] === 'x') {
      x2++
    } else {
      y2++
    }

    if (s1[i] == 'x' && s2[i] == 'y') {
      r1++
    }

    if (s1[i] == 'y' && s2[i] == 'x') {
      r2++
    }
  }

  if (x1 % 2 !== x2 % 2 || y1 % 2 !== y2 % 2) {
    return -1
  }

  //console.log(r1, r2)

  return (r1 >> 1) + (r2 >> 1) + r1 % 2 + r2 % 2
}

console.log(minimumSwap('xy', 'yx'))
console.log(minimumSwap('xx', 'yx'))
console.log(minimumSwap('xx', 'yy'))
console.log(minimumSwap('xxyyxyxyxx', 'xyyxyxxxyx'))
