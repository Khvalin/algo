class Solution:
    def getLucky(self, s: str, k: int) -> int:
        res = 0

        for ch in s:
            n = ord(ch) - ord("a") + 1
            while n > 0:
                res += n % 10
                n //= 10

        for _ in range(k - 1):
            t = 0
            while res > 0:
                t += res % 10
                res //= 10
            res = t

        return res


s = Solution()
print(s.getLucky("leetcode", 2))
