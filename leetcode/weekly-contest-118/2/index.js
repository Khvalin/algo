/**
 * @param {number[]} A
 * @return {number[]}
 */
var pancakeSort = function(A) {
  let res = []

  let flip = (index) => {
    res.push(index + 1)
    let i = 0
    while (index > i) {
      let t = A[i]
      A[i] = A[index]
      A[index] = t
      index--
      i++
    }
  }

  let sorted = [ ...A ].sort((a, b) => a - b)
  for (let i = A.length - 1; i >= 0; i--) {
    if (A[i] != sorted[i]) {
      let j = i
      while (A[j] != sorted[i]) {
        j--
      }

      if (j > 0) {
        flip(j)
      }
      flip(i)
    }
    // console.log(A)
  }

  return res
}

console.log(pancakeSort([ 3, 2, 4, 1 ]))
console.log(pancakeSort([ 3, 3, 3, 1 ]))
