/**
 * @param {number[]} A
 * @return {number[]}
 */
var sortArrayByParityII = function(A) {
  const res = Array.from(A)
  let oddI = 1,
    evenI = 0

  for (const n of A) {
    if (n % 2 === 0) {
      res[evenI] = n
      evenI += 2
    } else {
      res[oddI] = n
      oddI += 2
    }
  }

  return res
}
