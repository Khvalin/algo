/**
 * @param {number} n
 * @return {number}
 */
var countArrangement = function(n) {
    const solve = (nums, mask) => {
        let full = true
        for (const b of mask) {
            full = full && b
        }
        if (full || !nums.length) {
            return 1
        }
        
        let f = false
        let res = 0
        const n = nums[0]

        for (let i = 1; i <= mask.length;i++) {
            if (mask[i - 1] || (n % i > 0 && i % n >0) ) {
                continue
            }

            f = true
            const mCopy = [...mask]
            const numsCopy = [...nums]
            mCopy[i-1] = true
            numsCopy.shift()
            
            const t = solve(numsCopy, mCopy)
            if (!Number.isNaN(t)) {
                res+=t
            }
        }

        
        if (!f) {
            return Number.NaN
        }
        console.log(nums, mask, res)
        
        return res
    }
    
    const nums = []
    const mask = []
    for (let i = 1; i <= n; i++) {
        nums.push(i)
        mask.push(false)
    }
    
    return solve(nums, mask)
};


//console.log(countArrangement(2))
console.log(countArrangement(3))
console.log(countArrangement(4))
