from collections import Counter


class Solution:
    def frequencySort(self, s: str) -> str:
        count = Counter()
        for ch in s:
            count[ch] += 1

        chars = list(count.keys())
        chars.sort(key=count.get, reverse=True)

        return "".join([ch * count[ch] for ch in chars])


sol = Solution()
print(sol.frequencySort("Aabb"))
