/**
 * @param {number} n
 * @param {number} k
 * @return {number}
 */
 var kInversePairs = function(n, k) {
    const MOD = 1_000_000_000 + 7
    const memo = new Map()

    const solve = (n, k) => {
        if (k < 0 || n < 0) {
            return 0
        }
        
        if (n <= 1 && k > 0) {
            return 0
        }

        const total = (n * (n - 1)) >> 1
        if (k > total) {
            return 0
        }
        
        if (k === 0) {
            return 1
        }
        
        const key = n * 1_000_001 + k
        if (memo.has(key)) {
            return memo.get(key)
        }
        
        let res = 0
        for (let i = 1; i <= n; i++) {
            const c = n - i
            if (c > k) {
                continue
            }
            
            const t = solve(n - 1, k - c)
            
            res += t % MOD
            res %= MOD
        }
        
        memo.set(key, res)
        
        return res
    }
    
    return solve(n, k)
};

//console.log(kInversePairs(999,990))

// console.log(kInversePairs(2,0))
// console.log(kInversePairs(2,1))
// console.log()

// console.log(kInversePairs(3,0))
// console.log(kInversePairs(3,1))
// console.log(kInversePairs(3,2))
// console.log(kInversePairs(3,3))

// console.log(kInversePairs(4,0))
// console.log(kInversePairs(4,1))
// console.log(kInversePairs(4,2))
// console.log(kInversePairs(4,3))
// console.log(kInversePairs(4,4))
// console.log(kInversePairs(4,5))
// console.log(kInversePairs(4,6))
// console.log();

// console.log(kInversePairs(5,0))
// console.log(kInversePairs(5,1))
// console.log(kInversePairs(5,2))
// console.log(kInversePairs(5,3))
// console.log(kInversePairs(5,4))
// console.log(kInversePairs(5,5))
// console.log(kInversePairs(5,6))
// console.log(kInversePairs(5,7))
// console.log(kInversePairs(5,8))
// console.log(kInversePairs(5,9))
// console.log(kInversePairs(5,10))