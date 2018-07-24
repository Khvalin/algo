/**
 * @param {number[]} A
 * @return {number}
 */
var lenLongestFibSubseq = function(A, pivot) {
  let max = 0

  const map = {}
  result = A.map(() => 0)
  for (let i = 0; i < A.length; i++) {
    map[A[i]] = i
  }

  for (let i = 0; i < A.length - 3; i++) {
    for (let j = i + 1; j < A.length - 2; j++) {
      let first = i
      let second = j
      let third = 0
      let cur = 0

      const tmp = [ first, second ]
      while ((third = map[A[first] + A[second]])) {
        //console.log(first, second, third)
        //
        // console.log(A[first], A[second], A[third])
        tmp.push(third)

        cur++
        first = second
        second = third
      }

      if (cur > max) {
        // console.log(tmp.map((i) => A[i]))
        max = cur
      }
    }
  }

  return max === 0 ? max : max + 2
}

console.log(lenLongestFibSubseq([ 1, 2, 3, 4, 5, 6, 7, 8 ]))
console.log(lenLongestFibSubseq([ 1, 3, 7, 11, 12, 14, 18 ]))
console.log(lenLongestFibSubseq([ 2, 4, 7, 8, 9, 10, 14, 15, 18, 23, 32, 50 ]))
