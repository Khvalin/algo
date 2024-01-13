/**
 * @param {string} s
 * @param {number} k
 * @return {number}
 */
var numberOfArrays = function(s, k) {
	const MOD = 1e9 + 7
	const digits = [...s].map(ch => 1 * ch)
	const memo = digits.map(() => -1)
	
	const solve = (ind) => {
		if (ind >= digits.length) {
			return 1
		}
		
		if (digits[ind] === 0) {
			return 0
		}
		
		if (memo[ind] >= 0) {
			return memo[ind]
		}
		
		let n = 0
		let res = 0
		let i = ind
		do {
			n *= 10
			n += digits[i++]
			if (n <= k) {
				res += solve(i)
				res %= MOD
			}
		} while (n <= k && i < digits.length)
		
		memo[ind] = res
		
		return res
	}
	
	
	return solve(0)
};

/*
 * 1 1 1
 * 1, 11
 * 11, 1
 * 111
 * 
 * */


let res = numberOfArrays('787871',2000)
console.log(24, res)
/*
res = numberOfArrays('1317', 2000)
console.log(res)


res = numberOfArrays('111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111', 2000)
console.log(res)
*/
