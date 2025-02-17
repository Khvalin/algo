from functools import lru_cache


class Solution:

    def numTilePossibilities(self, tiles: str) -> int:

        @lru_cache(maxsize=None)  # Cache unlimited results
        def solve(tiles):
            res = set()
            if len(tiles) == 0:
                return res
            if len(tiles) == 1:
                res.add(tiles[0])
                return res

            for i in range(len(tiles)):
                t = list(tiles)
                ch = t[i]
                a = t[:i] + t[i + 1 :]
                sub_res = solve(tuple(a))
                for b in sub_res:
                    res.add(ch + b)
                    res.add(b)

            return res

        input = tuple(tiles)
        r = solve(input)

        return len(r)
