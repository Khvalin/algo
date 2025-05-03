from typing import List


class Solution:
    def minDominoRotations(self, tops: List[int], bottoms: List[int]) -> int:
        dominos = [bottoms, tops]
        pair = [bottoms[0], tops[0]]
        res = [True] * 2

        counts = [[0, 0], [0, 0]]

        for i in range(len(tops)):
            for j in range(2):
                res[j] = res[j] and (
                    pair[j] == dominos[0][i] or pair[j] == dominos[1][i]
                )
                if not res[j]:
                    continue
                if pair[j] != dominos[j][i]:
                    counts[j][j] += 1
                k = (j + 1) % 2
                if pair[j] != dominos[k][i]:
                    counts[j][k] += 1
        ans = len(tops) + 1
        for i in range(2):
            if res[i]:
                ans = min(ans, *counts[i])

        if ans > len(tops):
            return -1
        return ans
