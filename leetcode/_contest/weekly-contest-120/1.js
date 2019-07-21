/**
 * @param {number[]} A
 * @return {number[]}
 */
var sortedSquares = function(A) {
  let sqs = A.map((n) => n * n)
  sqs = sqs.sort((a, b) => a - b)
  return sqs
}
