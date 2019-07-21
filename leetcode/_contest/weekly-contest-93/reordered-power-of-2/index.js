/**
 * @param {number} N
 * @return {boolean}
 */
var reorderedPowerOf2 = function(N) {
  const N_STR = N.toString();
  const LEN = N.toString().length;
  
  const norm = (str) => {
    return str.split('').sort().join('');
  };

  let pow = 0;
  let res = 1;
  let resStr = '1';

  do {
    if (resStr.length === LEN) {
      if (norm(resStr) === norm(N_STR)) {
          //console.log(norm(resStr), norm(N_STR))
          return true;
      }
    }

    res = 2 ** ++pow;
    resStr = res.toString();

  } while (resStr.length <= LEN);

  return false;

};


console.log(reorderedPowerOf2(1))
console.log(reorderedPowerOf2(24))
console.log(reorderedPowerOf2(46))