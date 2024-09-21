from typing import List


class Solution:
    def lexicalOrder(self, n: int) -> List[int]:
        res = []
        cur = 1

        while len(res) < n:
            res.append(cur)
            if cur * 10 <= n:
                cur *= 10
                continue
            f = True
            if cur % 10 < 9:
                if cur + 1 <= n:
                    cur += 1
                    continue
                cur += 1
                while cur > n:
                    cur //= 10
                    cur += 1
                    f = False

            if f:
                cur += 1
            while cur % 10 == 0:
                cur //= 10

        return res


s = Solution()
print(s.lexicalOrder(192))

# sample = []
# for i in range(1, 1000):
#     sample.append(f"{i}")
# sample.sort()
# print(sample)
