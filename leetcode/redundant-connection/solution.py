from typing import List


class Solution:
    def findRedundantConnection(self, edges: List[List[int]]) -> List[int]:
        parent = [0] * (1 + len(edges))
        children = [None] * (1 + len(edges))

        for i in range(len(parent)):
            children[i] = set([i])
            parent[i] = i

        def union(u, v):
            pu = parent[u]
            pv = parent[v]

            minp = min(pu, pv)
            maxp = max(pu, pv)

            for j in children[maxp]:
                parent[j] = minp
            children[minp].update(children[maxp])
            children[maxp].clear()

        res = edges[0]
        for u, v in edges:
            if parent[u] == parent[v]:
                res = [u, v]

            union(u, v)

        return res


test = Solution()
# print(test.findRedundantConnection([[1,2],[2,3],[3,4],[1,4],[1,5]]))

#
print(test.findRedundantConnection([[1, 5], [3, 4], [3, 5], [4, 5], [2, 4]]))
