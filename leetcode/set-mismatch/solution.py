from typing import List


class Solution:
    def findErrorNums(self, nums: List[int]) -> List[int]:
        res = [len(nums) + 1, len(nums) + 1]
        for n in nums:
            key = abs(n) - 1
            if nums[key] < 0:
                res[0] = key + 1
            else:
                nums[key] = -nums[key]
        for i in range(len(nums) + 1):
            if i >= len(nums) or nums[i] > 0:
                res[1] = i + 1
                break
        return res
