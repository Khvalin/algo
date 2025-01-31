class Solution:
    def largestIsland(self, grid: List[List[int]]) -> int:
        cc = []
        visited = []

        def bfs(x, y, k):
            res = 0
            q = [[x, y]]
            while len(q) > 0:
                [x, y] = q.pop()
                if cc[x][y] > 0:
                    continue
                res += 1
                cc[x][y] = k
                cells = [[x - 1, y], [x + 1, y], [x, y - 1], [x, y + 1]]
                for [x1, y1] in cells:
                    if x1 < 0 or x1 >= len(grid) or y1 < 0 or y1 >= len(grid[x1]):
                        continue
                    if grid[x1][y1] == 0 or visited[x1][y1]:
                        continue
                    visited[x1][y1] = True
                    q.append([x1, y1])

            return res

        for row in grid:
            cc.append([0] * len(row))
            visited.append([False] * len(row))

        res = 0
        groups = [0]
        for i, row in enumerate(grid):
            for j, n in enumerate(row):
                if n == 1 and cc[i][j] == 0:
                    count = bfs(i, j, len(groups))
                    groups.append(count)
                    res = max(res, count)

        for i, row in enumerate(grid):
            for j, n in enumerate(row):
                if n != 0:
                    continue
                count = 1
                cells = [[i - 1, j], [i + 1, j], [i, j - 1], [i, j + 1]]

                islands = set()
                for [x, y] in cells:
                    if x < 0 or x >= len(grid) or y < 0 or y >= len(grid[x]):
                        continue
                    islands.add(cc[x][y])

                for k in islands:
                    count += groups[k]
                res = max(res, count)

        return res
