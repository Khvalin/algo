/**
 * @param {number} K
 * @return {number}
 */
var preimageSizeFZF = function(K) {
  function getBaseLog(x, y) {
    return Math.log(y) / Math.log(x);
  }

  let count = 5;

  for (let i = 0; i < K; i++) {
    if (i % 10 === 0) {
      count += getBaseLog(i, 10);
    } else {
      if (i % 5 === 0) {
        count += getBaseLog(i, 5);
      }
    }
  }

  return count;
};

console.log(preimageSizeFZF(0));
