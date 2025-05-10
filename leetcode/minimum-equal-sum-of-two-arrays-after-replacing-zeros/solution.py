from typing import List

class Solution:
    def minSum(self, nums1: List[int], nums2: List[int]) -> int:
        a1 = [0, 0]
        a2 = [0, 0]

        for n in nums1:
            a1[0] += n
            if n == 0:
                a1[1] += 1

        for n in nums2:
            a2[0] += n
            if n == 0:
                a2[1] += 1

        s1 = sum(a1)
        s2 = sum(a2)

        if s1 == s2:
            return s1

        if (
            (a1[1] > 0 and a2[1] > 0) or
            (a1[1] == 0 and a2[1] > 0 and s1 > s2) or
            (a2[1] == 0 and a1[1] > 0 and s2 > s1)
        ):
            return max(s1, s2)

        return -1
