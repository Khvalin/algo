/**
 * @param {number[]} A
 * @param {number[]} B
 * @return {number}
 */
var minDominoRotations = function(A, B) {
  let aCount = {}
  let bCount = {}
  let n = A.length
  let [ aCand, bCand ] = [ -1, -1 ]

  for (let i = 0; i < n; i++) {
    aCount[A[i]] = aCount[A[i]] || 0
    aCount[A[i]]++

    if (aCount[A[i]] >= n / 2) {
      aCand = A[i]
    }

    bCount[B[i]] = bCount[B[i]] || 0
    bCount[B[i]]++
    if (bCount[B[i]] >= n / 2) {
      bCand = B[i]
    }
  }

  /*  console.log(aCount)
  console.log(bCount)
  console.log(aCand, bCand) */

  if (aCand == -1 && bCand == -1) {
    return -1
  }

  let cand = aCand
  if (aCand == -1 || aCount[aCand] < bCount[bCand]) {
    cand = bCand
    let t = Object.assign({}, aCand)
    aCand = Object.assign({}, bCand)
    bCand = Object.assign({}, t)

    t = [ ...A ]
    A = [ ...B ]
    B = t
  }

  let r = 0
  for (let i = 0; i < n; i++) {
    if (A[i] !== cand) {
      if (B[i] == cand) {
        r++
      } else {
        return -1
      }
    }
  }

  return r
}

console.log(minDominoRotations([ 5, 1, 2, 4, 2, 2 ], [ 2, 2, 6, 2, 3, 2 ]))
console.log(minDominoRotations([ 3, 5, 1, 2, 3 ], [ 3, 6, 3, 3, 4 ]))
