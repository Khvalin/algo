class Solution:
    def robotWithString(self, s: str) -> str:
        count = [0] * (ord("z") + 1)
        for ch in s:
            count[ord(ch)] += 1

        res = []
        t = []

        i = 0
        j = 0
        while i < len(s):
            while count[ j] == 0:
                j += 1
            ch = chr(j)

            while len(t) and t[-1] <= ch:
                res.append(t.pop())

            count[ord(s[i])] -= 1
            t.append(s[i])
            i += 1

        while len(t):
            res.append(t.pop())

        return "".join(res)
