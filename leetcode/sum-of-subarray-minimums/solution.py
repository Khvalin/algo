from typing import List


class Solution:
    def sumSubarrayMins(self, arr: List[int]) -> int:
        mod = 10**9 + 7
        inf = max(arr) + 1
        def solve(left, right):
            
            if left >= right:
                return None
            if left == right - 1:
                return [arr[left], arr[left]]

            m = inf
            s = 0
            mid = (left + right) // 2
            left_result = solve(left, mid)

            m = min(m, left_result[0])
            s = left_result[1]
            right_result = solve(mid, right)

            m = min(m, right_result[0])
            s += right_result[1]
            s += m
            s %= mod
            print(arr[left:right])
            return [m, s]
            
        return solve(0, len(arr))[1]
            
print(Solution().sumSubarrayMins([3,1,2,4]))
#print(Solution().sumSubarrayMins([1,1,1,1]))
