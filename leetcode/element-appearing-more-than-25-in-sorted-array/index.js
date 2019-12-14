/**
 * @param {number[]} arr
 * @return {number}
 */
var findSpecialInteger = function(arr) {
  let t = null;
  let c = 0;
  for (const n of arr) {
    if (n === t) {
      c++;
    } else {
      if (c > arr.length >> 2) {
        return t;
      }
      c = 0;
      t = n;
    }
  }

  return arr.slice(-1);
};
