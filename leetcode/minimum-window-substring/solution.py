class Solution:
    def minWindow(self, s: str, t: str) -> str:
        t_counts = {}
        for ch in t:
            t_counts[ch] = t_counts.get(ch, 0) + 1

        s_counts = {}
        res = (-1, len(s) + 1)
        c = 0
        l = 0
        for i in range(0, len(s)):
            ch = s[i]
            if ch not in t_counts:
                continue
            s_counts[ch] = s_counts.get(ch, 0) + 1
            if s_counts[ch] == t_counts[ch]:
                c += 1
            if c < len(t_counts):
                continue

            while s[l] not in t_counts or s_counts[s[l]] > t_counts[s[l]]:
                if s[l] in s_counts:
                    s_counts[s[l]] -= 1
                l += 1
            if i + 1 - l < res[1] - res[0]:
                res = (l, i + 1)

        if res[0] < 0:
            return ""
        return s[res[0] : res[1]]


solution = Solution()
res = solution.minWindow("ADOBECODEBANC", "ABC")
print(res)

res = solution.minWindow("bba", "ab")
print(res)
