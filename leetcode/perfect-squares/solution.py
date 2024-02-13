import math

class Solution:
    def numSquares(self, n: int) -> int:
        res = [n + 3] * (n + 1)
        res[0] = 0
        for i in range(1, n+1):
            for j in range(1+math.floor(math.sqrt(i)), 0, -1):
                k = i - j*j
                if k < 0:
                    continue
                res[i] = min(res[i], res[k] + 1)
        return res[n]
