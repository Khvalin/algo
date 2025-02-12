class Solution:
    def removeOccurrences(self, s: str, part: str) -> str:
        prefix = [None] * len(s)
        remove = [False] * len(s)
        prev = []

        for i in range(len(s)):
            ch = s[i]
            a = list(prev)
            a.append(0)
            nextp = []
            for j in a:
                if part[j] != ch:
                    continue
                j += 1
                if j == len(part):
                    cnt = len(part)
                    last = i
                    while cnt > 0:
                        if not remove[last]:
                            remove[last] = True
                            cnt -= 1
                        last -= 1

                    print(i, last, prev)
                    prev = prefix[last] if last >= 0 else []
                    break
                else:
                    nextp.append(j)
            if remove[i]:
                continue
            prefix[i] = list(nextp)
            prev = list(nextp)

        res = []
        for i in range(len(s)):
            if remove[i]:
                continue
            res.append(s[i])
        return "".join(res)


s = "daabcbaabcbc"
part = "abc"

s = "hhvhvaahvahvhvaavhvaasshvahvaln"
part = "hva"

print(Solution().removeOccurrences(s, part))
