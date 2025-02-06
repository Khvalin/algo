class Solution:
    def tupleSameProduct(self, nums: List[int]) -> int:
        count = {}
        for i in range(len(nums)):
            for j in range(i + 1, len(nums)):
                n = nums[i] * nums[j]
                c = count.get(n) or 0
                count[n] = c + 1
        res = 0
        for v in count.values():
            if v > 1:
                res += 2 * math.comb(v, 2) * 4

        return res
