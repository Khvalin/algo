from typing import List
from functools import cache


class Solution:
    def diffWaysToCompute(self, expression: str) -> List[int]:

        # @cache
        def solve(expression):
            print(expression)
            a = list(expression)
            if len(a) == 0:
                return []

            if len(a) <= 1:
                return a

            res = []

            p = 0
            match a[1]:
                case "-":
                    p = a[0] - a[2]
                case "+":
                    p = a[0] + a[2]
                case "*":
                    p = a[0] * a[2]
                case _:
                    raise ValueError(a[i])

            clone = [p] + a[3:]
            res += solve(tuple(clone))

            return res

        parsed = [0]
        for i, ch in enumerate(expression):
            match ch:
                case "-":
                    parsed.append("-")
                    parsed.append(0)
                case "+":
                    parsed.append("+")
                    parsed.append(0)
                case "*":
                    parsed.append("*")
                    parsed.append(0)
                case _:
                    parsed[-1] *= 10
                    parsed[-1] += ord(ch) - ord("0")

        res = solve(tuple(parsed))
        return res


s = Solution()
# print(s.diffWaysToCompute("2-1-1"), [0,2])
# print(s.diffWaysToCompute("2*3-4*5"), [-34,-14,-10,-10,10])
print(s.diffWaysToCompute("2*2*2*2"), [16, 16, 16, 16, 16])
