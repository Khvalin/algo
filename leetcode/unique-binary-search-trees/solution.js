/**
 * @param {number} n
 * @return {number}
 */
 var numTrees = function(n) {
  const memo = new Map()
  memo.set(0, 1)
  memo.set(1, 1)
  
  const count = (n) => {
      if (n < 2) {
          return 1
      }
      if (memo.has(n)) {
          return memo.get(n)
      }
      let res = 0
      for (let i = 0; i < n; i++) {
          const left = count(i)
          const right = count(n - i-1)
          
          res += left * right
      }
       memo.set(n, res)
      
      return res
  }
  
  
  return count(n)
};
