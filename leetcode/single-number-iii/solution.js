/**
 * @param {number[]} nums
 * @return {number[]}
 */
 var singleNumber = function(nums) {
  let xor = 0
  for (const n of nums) {
      xor ^= n
  }
  
  let mask = 1;
  while (!(xor & mask)) {
      mask = mask << 1;
  }
  
  let [a, b] = [0, 0]
  for (const n of nums) {
      if (n & mask) {
          a ^= n
          continue
      }
      b ^= n
  }
  
  
  return [a, b]
};