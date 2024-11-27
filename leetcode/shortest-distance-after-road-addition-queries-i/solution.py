class Solution:
    def shortestDistanceAfterQueries(
        self, n: int, queries: List[List[int]]
    ) -> List[int]:
        dist = [i for i in range(n)]
        adj = [[]]
        for i in range(1, n):
            adj.append([i - 1])

        res = []
        for q in queries:
            [u, v] = q

            adj[v].append(u)

            for i in range(v, n):
                for j in adj[i]:
                    dist[i] = min(dist[i], dist[j] + 1)

            res.append(dist[-1])

        return res
