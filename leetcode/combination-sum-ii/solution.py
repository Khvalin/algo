from typing import List

# from collections import defaultdict


class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        sums = [[], []]
        candidates.sort()
        for _ in range(target + 1):
            sums[0].append(set())
            sums[1].append(set())

        cur = 0
        l = 0
        for i in range(0, len(candidates)):
            n = candidates[i]
            if n > target:
                continue

            cur = l % 2
            prev = (l + 1) % 2
            l += 1
            # for a in sums[cur]:
            # a.clear()
            sums[cur][target - n].add((n,))
            for k, arr in enumerate(sums[prev]):
                for a in arr:
                    sums[cur][k].add(a)
                    if k < n:
                        continue

                    sums[cur][k - n].add(a + (n,))

        return list(sums[cur][0])


t1 = Solution().combinationSum2([10, 1, 2, 7, 6, 1, 5], target=8)
print(t1)
