/**
 * @param {number[]} heights
 * @return {number}
 */
var heightChecker = function(heights) {
  let sorted = [ ...heights ].sort((a, b) => a - b)
  let r = 0
  for (let i = 0; i < sorted.length; i++) {
    if (sorted[i] != heights[i]) {
      r++
    }
  }

  return r
}
