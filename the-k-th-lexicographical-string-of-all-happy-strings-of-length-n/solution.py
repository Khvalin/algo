class Solution:
    def getHappyString(self, n: int, k: int) -> str:
        N = 3
        self.a = [i % 2 for i in range(n)]

        def inc(p):
            self.a[p] += 1

            if self.a[p] >= N:
                if p <= 0:
                    self.a = []
                    return
                return inc(p - 1)
            if p > 0 and self.a[p - 1] == self.a[p]:
                return inc(p)

            for i in range(p + 1, n):
                self.a[i] = 0 if self.a[i - 1] > 0 else 1

        for i in range(1, k):
            inc(n - 1)
            if len(self.a) == 0:
                break

        return "".join([["a", "b", "c"][i] for i in self.a])
