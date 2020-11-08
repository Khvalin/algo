/**
 * @param {number} poured
 * @param {number} query_row
 * @param {number} query_glass
 * @return {number}
 */
var champagneTower = function(poured, query_row, query_glass) {
  const memo = new Map()

  const solve = (poured, query_row, query_glass)  => {
    if (poured < Number.EPSILON) {
      return 0
    }
    if (query_row < 0 || query_glass < 0 || query_glass > query_row) {
      return 0
    }

    if (query_row === 0) {
      return poured
    }
    
    const key = 100 * query_row + query_glass
    if (memo.has(key)) {
      return memo.get(key)
    }
    
    const p1 = Math.max(solve(poured, query_row - 1, query_glass-1)-1 , 0)
    const p2 = Math.max(solve(poured, query_row - 1, query_glass)-1, 0)
    
    const res = (p1 + p2) / 2
    memo.set(key, res)
    return res
  }
  
 return Math.min(solve(poured, query_row, query_glass), 1)
};


console.log(champagneTower(50,5,5))
console.log(champagneTower(100000009,33,17))
