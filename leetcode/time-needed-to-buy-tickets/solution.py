class Solution:
    def timeRequiredToBuy(self, tickets: List[int], k: int) -> int:
        res = 0

        t = tickets[k]

        for i in range(0, len(tickets)):
            d = 0
            if (i > k):
                d = 1
            res += min(tickets[i], t - d)

        return res
