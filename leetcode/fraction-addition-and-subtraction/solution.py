from math import gcd


class Solution:
    def fractionAddition(self, expression: str) -> str:
        fractions = [[1, 0, 0]]

        pos = 1
        for i, ch in enumerate(expression):
            match ch:
                case "+":
                    pos = 1
                    fractions.append([1, 0, 0])
                case "-":
                    if i == 0:
                        fractions.pop()
                    pos = 1
                    fractions.append([-1, 0, 0])
                case "/":
                    pos = 2
                case _:
                    fractions[-1][pos] *= 10
                    fractions[-1][pos] += ord(ch) - ord("0")

        # def gcd(a, b):
        #     if b == 0:
        #         return a
        #     return gcd(b, a%b)

        def simplify(n, m):
            if n == 0:
                return (0, 1)
            g = gcd(n, m)
            n //= g
            m //= g

            return (n, m)

        res = [1, 0, 1]
        for i, cur in enumerate(fractions):
            if i == 0:
                res = cur
                continue

            n = res[0] * res[1] * cur[2] + cur[0] * cur[1] * res[2]
            m = res[2] * cur[2]

            d = 1
            if n < 0:
                d = -1
                n = -n

            n, m = simplify(n, m)
            res[0] = d
            res[1] = n
            res[2] = m

        d, n, m = res
        return f"{'-' if d < 0 else '' }{n}/{m}"


sol = Solution()
# print(sol.fractionAddition("-1/2+1/2"))
# print(sol.fractionAddition("-1/2+1/2+1/3"))
print(sol.fractionAddition("5/3+1/3"))
