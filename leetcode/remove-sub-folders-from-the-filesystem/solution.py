class Solution:
    def removeSubfolders(self, folder: List[str]) -> List[str]:
        trie = {}
        for f in folder:
            cur = trie
            for fragment in f.split("/"):
                if fragment not in cur:
                    cur[fragment] = {}
                cur = cur[fragment]
            cur["!"] = True
        res = set()

        def walk(node, path):
            if node.get("!", False):
                res.add("/".join(path))
                return

            for key, value in node.items():
                c = path.copy()
                c.append(key)
                walk(value, c)

        walk(trie, [])

        return list(res)
