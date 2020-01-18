/**
 * @param {number[]} nums
 * @param {number} S
 * @return {number}
 */
const findTargetSumWays = (nums, S) => {
    const memo = [new Map(), new Map()]
    memo[1].set(0, 1)

    for (let i = 0; i < nums.length; i++){
        const [cur, prev] = [i%2, (i+1)%2]
        memo[cur].clear()
        for (const [key, val] of memo[prev]) {
            for (let k = -1; k <= 1; k += 2){
                const count = memo[cur].get(key + k*nums[i]) || 0
                memo[cur].set(key + k*nums[i], val + count)
            }
        }
    }

    return memo[(nums.length + 1) % 2].get(S) || 0
};


console.log(findTargetSumWays([1, 1, 1, 1, 1], 3 ))
