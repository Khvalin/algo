/**
 * @param {number} X
 * @param {number} Y
 * @return {number}
 */
var brokenCalc = function(x, y) {
  let res = 0;
  while (x < y) {
    if (y % 2 === 1) {
      res++;
      y++;
      continue;
    }

    res++;
    y /= 2;
  }

  return res + x - y;
};

// console.log(brokenCalc(2, 3));
// console.log(brokenCalc(3, 10));
console.log(brokenCalc(3, 15110));
