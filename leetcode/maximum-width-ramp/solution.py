from typing import List


class Solution:
    def maxWidthRamp(self, nums: List[int]) -> int:
        stack = []
        for i in range(len(nums)):
            if not len(stack):
                stack.append(i)
                continue
            if nums[i] < nums[stack[-1]]:
                stack.append(i)
        res = 0

        for i in reversed(range(len(nums))):
            while len(stack) and nums[i] >= nums[stack[-1]]:
                j = stack.pop()
                res = max(res, i - j)
            if not len(stack):
                break

        return res


s = Solution()
print(s.maxWidthRamp([6, 0, 8, 2, 1, 5]), 4)
print(s.maxWidthRamp([9, 8, 1, 0, 1, 9, 4, 0, 4, 1]), 7)
