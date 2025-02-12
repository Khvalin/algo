from typing import List


class Solution:
    def maximumSum(self, nums: List[int]) -> int:
        res = -1
        sums = {}
        for n in nums:
            s = 0
            t = n
            while t > 0:
                s += t % 10
                t //= 10

            if s not in sums:
                sums[s] = []

            sums[s].append(n)

            sums[s].sort(reverse=True)
            if len(sums[s]) > 2:
                sums[s].pop()

            if len(sums[s]) > 1:
                res = max(res, sums[s][0] + sums[s][1])

        return res
