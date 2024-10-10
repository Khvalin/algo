/**
 * @param {number[]} nums
 * @return {number}
 */
var maxWidthRamp = function(nums) {
    const items = nums.map((n, i) => [n, i])
    items.sort((a, b) => a[0] - b[0])
    
    let res = 0
    let i = items[0][1]
    for (const [_, j] of items) {
		if (j > i) {
			res = Math.max(res, j - i)
		}
		i = Math.min(i, j)
	}
    return res
};

console.log(maxWidthRamp([9,8,1,0,1,9,4,0,4,1]))
