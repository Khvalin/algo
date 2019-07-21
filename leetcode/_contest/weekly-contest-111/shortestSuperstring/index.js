const U = {}
const concat = (s1, s2) => {
  return s1 + '¿' + s2
}

const getOverlap = (s1, s2) => {
  let i = s2.length + 1

  let res = null
  while (i > 0) {
    if (s1.endsWith(s2.slice(0, i))) {
      res = s1 + s2.slice(i)
      break
    }
    i--
  }
  if (!res) {
    res = s1 + s2
  }

  return {
    overlap: i,
    res
  }
}
/* 
  console.log(getOverlap('haha', 'hehe'))
  console.log(getOverlap('geek', 'eke'))
  console.log(getOverlap('ttca', 'ca tgcatc'))
  return
 */

/**
 * @param {string[]} A
 * @return {string}
 */
var shortestSuperstring = function(A) {
  if (A.length <= 1) {
    return A[0]
  }

  let overlaps = []
  let maxOverlap = -1

  for (let i = 0; i < A.length - 1; i++) {
    for (let j = i + 1; j < A.length; j++) {
      let con = concat(A[i], A[j])
      let u = U[con]
      if (!u) {
        //console.log(A[i], A[j])
        const u1 = getOverlap(A[i], A[j])
        const u2 = getOverlap(A[j], A[i])

        if (u1.overlap > u2.overlap) {
          u = u1
        } else {
          u = u2
        }
        U[con] = u
      }

      overlaps.push({ ...u, x: i, y: j })

      if (maxOverlap < u.overlap) {
        maxOverlap = u.overlap
      }
    }
  }

  let minLen = Number.POSITIVE_INFINITY
  let res = ''
  for (o of overlaps) {
    if (o.overlap === maxOverlap) {
      //    console.log(A[o.x], A[o.y], o)
      let copy = [ ...A ]
      copy[o.y] = copy[copy.length - 1]
      copy[o.x] = o.res
      copy.pop()

      let tmp = shortestSuperstring(copy)
      if (tmp.length < minLen) {
        minLen = tmp.length
        res = tmp
      }
    }
  }
  return res
  //    console.log(A)
}

//console.log(shortestSuperstring([ 'geek', 'eke' ]))
//console.log(shortestSuperstring([ 'alex', 'loves', 'leetcode' ]))
//console.log(shortestSuperstring([ 'catg', 'ctaagt', 'gcta', 'ttca', 'atgcatc' ]))
//console.log(shortestSuperstring([ 'ift', 'efd', 'dnete', 'tef', 'fdn' ]), 'iftefdnete')
console.log(shortestSuperstring([ 'mah', 'obr', 'ahkh', 'royy', 'broy' ]), 'obroyymahkh')
