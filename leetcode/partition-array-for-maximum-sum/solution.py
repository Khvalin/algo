class Solution:
    def maxSumAfterPartitioning(self, arr: List[int], k: int) -> int:

        memo = {}

        def solve(ind, m, c):
            if ind >= len(arr):
                return c * m

            key = (ind, c)
            if key in memo:
                return memo[key]

            m = max(m, arr[ind])
            c += 1
            res = c * m + solve(ind + 1, 0, 0)
            if c < k:
                res = max(res, solve(ind + 1, m, c))

            memo[key] = res
            return res

        return solve(0, 0, 0)
