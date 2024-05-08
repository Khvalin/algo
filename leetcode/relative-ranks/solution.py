from typing import List


class Solution:
    def findRelativeRanks(self, score: List[int]) -> List[str]:
        a = sorted(score, reverse=True)
        rank = {}
        for i, n in enumerate(a):
            rank[n] = i
        res = []
        medals = ["Gold Medal", "Silver Medal", "Bronze Medal"]
        for n in score:
            n = rank[n]
            r = f"{n+1}"
            if n < len(medals):
                r = medals[n]
            res.append(r)
        return res


res = Solution().findRelativeRanks([5, 4, 3, 2, 1])
print(res)
