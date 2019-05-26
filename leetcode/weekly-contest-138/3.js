/**
 * @param {number[]} A
 * @return {number[]}
 */
var prevPermOpt1 = function(A) {
  let min = A[A.length - 1]

  for (let i = A.length - 1; i >= 0; i--) {
    let a = A[i]

    let swapIndex = -1
    let m = -1
    for (let j = i + 1; j < A.length; j++) {
      if (a > A[j]) {
        if (A[j] >= m) {
          m = A[j]
          swapIndex = j
        }
      }
    }

    if (swapIndex > -1) {
      A[i] = m
      A[swapIndex] = a
      break
    }
  }

  return A
}

let ans = prevPermOpt1([ 3, 2, 1 ])
console.log(ans)

ans = prevPermOpt1([ 3, 1, 1, 3 ])
console.log(ans)

ans = prevPermOpt1([ 1, 9, 4, 6, 7 ])
console.log(ans)
/*
ans = prevPermOpt1([
  1,
  9,
  6,
  7,
  9,
  6,
  4,
  4,
  2,
  2,
  7,
  7,
  7,
  6,
  3,
  5,
  7,
  7,
  3,
  8,
  8,
  4,
  4,
  1,
  5,
  4,
  7,
  4,
  7,
  3,
  7,
  5,
  4,
  1,
  7,
  4,
  9,
  6,
  5,
  9,
  8,
  9,
  9,
  4,
  6,
  6,
  5,
  5,
  7,
  7,
  8,
  1,
  4,
  6,
  4,
  5,
  4,
  4,
  8,
  9,
  5,
  7,
  2,
  4
])
console.log(ans)
 */
