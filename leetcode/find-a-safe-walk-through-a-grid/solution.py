import heapq


class Solution:
    def findSafeWalk(self, grid: List[List[int]], health: int) -> bool:
        max_health = [[-10] * len(row) for row in grid]

        max_health[0][0] = health - grid[0][0]

        q = [(-max_health[0][0], 0, 0)]

        while q:
            h, x, y = heapq.heappop(q)
            h = -h

            if h < 1:
                continue

            if x == len(grid) - 1 and y == len(grid[x]) - 1:
                return True

            directions = [
                [-1, 0],
                [1, 0],
                [0, -1],
                [0, 1],
            ]

            for [dx, dy] in directions:
                x1 = x + dx
                y1 = y + dy

                if x1 < 0 or x1 >= len(grid) or y1 < 0 or y1 >= len(grid[x]):
                    continue

                cell_health = h - grid[x1][y1]
                if cell_health > 0 and max_health[x1][y1] < cell_health:
                    max_health[x1][y1] = cell_health
                    heapq.heappush(q, (-cell_health, x1, y1))

        return False
