class Solution:
    def countSubIslands(self, grid1: List[List[int]], grid2: List[List[int]]) -> int:
        visited = []
        for x in range(len(grid2)):
            visited.append([False] * len(grid2[x]))

        def find_island(x, y):
            res = True
            q = [(x, y)]
            dirs = [[0, 1], [0, -1], [-1, 0], [1, 0]]
            while len(q):
                x, y = q.pop()
                if not grid2[x][y]:
                    continue
                res = res and bool(grid1[x][y])
                for d in dirs:
                    dx = x + d[0]
                    dy = y + d[1]
                    if dx < 0 or dy < 0 or dx >= len(grid2) or dy >= len(grid2[x]):
                        continue
                    if not grid2[dx][dy] or visited[dx][dy]:
                        continue
                    q.append((dx, dy))
                    visited[dx][dy] = True

            return res

        res = 0
        for x in range(len(grid2)):
            for y in range(len(grid2[x])):
                if grid2[x][y] and not visited[x][y]:
                    found = find_island(x, y)
                    if found:
                        res += 1

        return res
