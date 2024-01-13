/**
 * @param {number[]} nums
 * @return {number[]}
 */
var sortArray = function (nums) {
    const mergeSort = (l, r) => {
        if (l === r) {
            return [nums[l]]
        }
        const mid = (l + r) >> 1

        const left = mergeSort(l, mid)
        const right = mergeSort(mid + 1, r)

        const res = []

        let i = 0
        let j = 0

        while (i < left.length || j < right.length) {
            if (i >= left.length) {
                res.push(right[j++])
                continue
            }

            if (j >= right.length || left[i] < right[j]) {
                res.push(left[i++])
                continue
            }

            res.push(right[j++])
        }
        return res
    }

    const quickSort = (l, r) => {
        if (l >= r - 1) {
            return
        }

        function partition(left, right) {
            const pivotIndex = Math.floor(Math.random() * (right - left + 1) + left);
            const pivot = nums[pivotIndex];
            let i = left;

            [nums[pivotIndex], nums[right]] = [nums[right], nums[pivotIndex]];

            for (let j = left; j < right; j++) {
                if (nums[j] < pivot) {
                    [nums[i], nums[j]] = [nums[j], nums[i]];
                    i++;
                }
            }

            [nums[i], nums[right]] = [nums[right], nums[i]];
            return i;
        }


        const pivot = partition(l, r)

        quickSort(l, pivot - 1)
        quickSort(pivot + 1, r)
    }

    quickSort(0, nums.length - 1)
    return nums
};


console.log(sortArray([0, 1, 0, 1, 2, 5]))