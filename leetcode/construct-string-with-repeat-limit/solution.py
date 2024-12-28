class Solution:
    def repeatLimitedString(self, s: str, repeatLimit: int) -> str:
        cnt = [0] * 28

        for ch in s:
            cnt[ord("z") - ord(ch)] += 1

        res = [""] * len(s)
        cur = 0
        k = 0
        for j in range(len(s)):
            for i in range(len(cnt)):
                if cnt[i] <= 0:
                    continue
                ch = chr(ord("z") - i)
                if j > 0 and (res[j - 1] == ch and cur >= repeatLimit):
                    continue

                if j == 0 or res[j - 1] != ch:
                    cur = 0

                cnt[i] -= 1
                res[j] = ch
                cur += 1
                break
            else:
                break

        return "".join(res)
