class Solution:
    def countPalindromicSubsequence(self, s: str) -> int:
        seen = set()
        last = dict()

        for i, ch in enumerate(s):
            last[ch] = i

        res = 0
        for i, ch in enumerate(s):
            if ch in seen:
                continue
            seen.add(ch)

            a = set()
            for k in range(i + 1, last[ch]):
                a.add(s[k])
            res += len(a)

        return res


s = Solution()
print(s.countPalindromicSubsequence("aabcccca"))
