from functools import cache


class Solution:
    def strangePrinter(self, s: str) -> int:
        expected = s.encode("ascii")

        def create_copy(a, l, r, ch):
            copy = list(b"-" * len(a))
            for j in range(l, r):
                copy[j] = ch if ch is not None else a[j]
            return copy

        @cache
        def solve(a, l, r):
            # print(a, l)
            if l >= r or l >= len(expected) or r <= 0:
                return 0

            while l < r and a[l] == expected[l]:
                l += 1
            while l < r and a[r - 1] == expected[r - 1]:
                r -= 1
            if l >= r:
                return 0

            ch = expected[l]

            res = len(expected) + 10
            i = r
            while i >= l:
                if i == r or expected[i] == ch:
                    j = i
                    while i > l and expected[i - 1] == ch:
                        i -= 1
                    c1 = create_copy(a, l, i, ch)
                    c2 = create_copy(a, i + 1, r, None)
                    res = min(res, solve(bytes(c1), l, i) + solve(bytes(c2), j + 1, r))

                i -= 1

            return 1 + res

        start = bytes(create_copy(s, 0, len(s), expected[0]))

        return 1 + solve(start, 0, len(s))


s = Solution()
res = s.strangePrinter("abcdcbaz")
print(res, 5)
# res = s.strangePrinter("abac")
# print(res, 3)
res = s.strangePrinter("baacdddaaddaaaaccbddbcabdaabdbbcdcbbbacbddcabcaaa")
print(res, 19)

# res = s.strangePrinter("abababac")
# print(res, 5)
