/**
 * @param {number[]} A
 * @param {number[][]} queries
 * @return {number[]}
 */
var sumEvenAfterQueries = function(A, queries) {
  const res = []

  for (const q of queries) {
    A[q[1]] += q[0]

    let s = 0
    for (const n of A) {
      if (n % 2 == 0) {
        s += n
      }
    }

    res.push(s)
  }

  return A
}
