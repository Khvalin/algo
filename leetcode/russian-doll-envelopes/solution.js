import { input10 } from "./input.js";

/* returns the length of LIS in nums */
const lengthOfLIS = (nums) => {
    let piles = 0, n = nums.length;
    let top = nums.map(() => 0)
    for (let i = 0; i < n; i++) {
        // playing card to process
        const poker = nums[i]
        let [left, right ] = [0,piles] 
        // position to insert for binary search
        while (left < right) {
            let mid = (left + right) >> 1
            if (top[mid] >= poker) {
                right = mid
                continue
            }

            left = mid + 1;
        }

        if (left === piles) {
            piles++
        }
        // put this playing cart on top of the pile
        top[left] = poker;
    }
    // the number of cards is the LIS length
    return piles;
}

/**
 * @param {number[][]} envelopes
 * @return {number}
 */
var maxEnvelopes = function (envelopes) {
    envelopes = envelopes.sort((a, b) => {
        // sort by ascending width, and sort by descending height if the width are the same
        return a[0] == b[0] ?
            b[1] - a[1] : a[0] - b[0];
    })

    const heights = envelopes.map(a => a[1])

    return lengthOfLIS(heights)
};

console.log(maxEnvelopes(input10), 593)

