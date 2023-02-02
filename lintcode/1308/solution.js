export class Solution {
  /**
   * @param n: a integer
   * @return: return a 2D array
   *          we will sort your return value in output
   */
  getFactors(n) {
    const factors = new Map();
    
    const solve = (n) => {
		if (factors.has(n)) {
			return factors.get(n)
		}
		const res = []
		const max = Math.floor(Math.sqrt(n))
		
		for (let i = 2; i <= max; i++) {
			if (n % i !== 0) {
				continue
			}
			
			const factors = solve(n / i)
			for (const a of factors) {
				if (a[0] >= i) {
					res.push([i, ...a])
				}
			}
			res.push([i, n/i])
		}
		
		factors.set(n, res)
		
		return res
	}
	
	return solve(n)
  }
}


const s = new Solution()
console.log(s.getFactors(256))
