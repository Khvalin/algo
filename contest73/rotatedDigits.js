/**
 * @param {number} N
 * @return {number}
 */
var rotatedDigits = function(N) {
  const isGood = function(n) {
    const str = n.toString();
    //0, 1, and 8 rotate to themselves; 2 and 5 rotate to each other; 6 and 9 rotate to each other, and the rest of the numbers do not rotate to any other number.

    return (
      ((str.indexOf('3') === -1 && str.indexOf('4')) === -1 && str.indexOf('7')) === -1 &&
      (str.indexOf('2') > -1 || str.indexOf('5') > -1 || str.indexOf('6') > -1 || str.indexOf('9') > -1)
    );
  };

  let result = 0;
  for (let i = 2; i <= N; i++) {
    if (isGood(i)) {
      result += 1;
    }
  }

  return result;
};

console.log(rotatedDigits(10));
